package ghostty

import (
	"bytes"
	"fmt"
)

type restoreData struct {
	BundleID string
	Windows  []restoreWindow
}

type restoreWindow struct {
	SelectedTabIndex int
	Tabs             []restoreTab
}

type restoreTab struct {
	Index     int
	Name      string
	Terminals []restoreTerminal
	Splits    []splitOp
	Hints     []hintOp
}

type restoreTerminal struct {
	Name             string
	WorkingDirectory string
}

type splitOp struct {
	ParentVar        string
	NewVar           string
	CfgVar           string
	Direction        string
	WorkingDirectory string
}

type hintOp struct {
	Var  string
	Text string
}

func RestoreSession(session Session) error {
	script, err := generateRestoreScript(session)
	if err != nil {
		return err
	}

	_, err = execScript(script)
	return err
}

func RenderRestoreScript(session Session) (string, error) {
	return generateRestoreScript(session)
}

func generateRestoreScript(session Session) (string, error) {
	data := restoreData{BundleID: BundleID}

	for _, win := range session.Windows {
		rw := restoreWindow{}
		var selectedOriginalID string
		for _, tab := range win.Tabs {
			if tab.Selected {
				selectedOriginalID = tab.ID
			}
		}

		for _, tab := range win.Tabs {
			if len(tab.Terminals) == 0 {
				continue
			}

			rt := restoreTab{
				Index: len(rw.Tabs) + 1,
				Name:  tab.Name,
			}

			if tab.ID == selectedOriginalID {
				rw.SelectedTabIndex = rt.Index
			}

			for _, t := range tab.Terminals {
				rt.Terminals = append(rt.Terminals, restoreTerminal{
					Name:             t.Name,
					WorkingDirectory: t.WorkingDirectory,
				})
			}

			varMap := map[int]string{0: "t1"}
			if len(tab.Terminals) > 1 {
				bsp := &bspBuilder{
					terminals: tab.Terminals,
					nextVar:   2,
					varMap:    varMap,
				}

				bsp.build(0, len(tab.Terminals), "t1", true)
				rt.Splits = bsp.ops
				varMap = bsp.varMap
			}

			for i, t := range tab.Terminals {
				v := varMap[i]
				rt.Hints = append(rt.Hints, hintOp{
					Var:  v,
					Text: fmt.Sprintf("# [boo] was: %s", t.Name),
				})
			}

			rw.Tabs = append(rw.Tabs, rt)
		}

		if len(rw.Tabs) > 0 {
			data.Windows = append(data.Windows, rw)
		}
	}

	var buf bytes.Buffer
	if err := templates.ExecuteTemplate(&buf, "session_restore.applescript", data); err != nil {
		return "", fmt.Errorf("render restore script: %w", err)
	}

	return buf.String(), nil
}

type bspBuilder struct {
	terminals []SessionTerminal
	nextVar   int
	varMap    map[int]string
	ops       []splitOp
}

func (s *bspBuilder) build(lo, hi int, parentVar string, horizontal bool) {
	count := hi - lo
	if count <= 1 {
		return
	}

	leftCount := (count + 1) / 2
	rightStart := lo + leftCount

	dir := "right"
	if !horizontal {
		dir = "down"
	}

	newVar := fmt.Sprintf("t%d", s.nextVar)
	cfgVar := fmt.Sprintf("cfg_%s", newVar)
	s.nextVar++

	s.varMap[rightStart] = newVar
	s.ops = append(s.ops, splitOp{
		ParentVar:        parentVar,
		NewVar:           newVar,
		CfgVar:           cfgVar,
		Direction:        dir,
		WorkingDirectory: s.terminals[rightStart].WorkingDirectory,
	})

	s.build(lo, rightStart, parentVar, !horizontal)
	s.build(rightStart, hi, newVar, !horizontal)
}
