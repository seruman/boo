package ghostty

import "fmt"

func CaptureSession() (Session, error) {
	windows, err := ListWindows()
	if err != nil {
		return Session{}, err
	}

	var session Session
	for _, win := range windows {
		sw := SessionWindow{
			ID:            win.ID,
			Name:          win.Name,
			SelectedTabID: win.SelectedTabID,
		}

		tabs, err := ListTabs(win.ID)
		if err != nil {
			return Session{}, fmt.Errorf("listing tabs for window %s: %w", win.ID, err)
		}

		for _, tab := range tabs {
			st := SessionTab{
				ID:       tab.ID,
				Name:     tab.Name,
				Index:    tab.Index,
				Selected: tab.Selected,
			}

			terminals, err := ListTerminalsOfTab(win.ID, tab.ID)
			if err != nil {
				return Session{}, fmt.Errorf("listing terminals for tab %s: %w", tab.ID, err)
			}

			for _, term := range terminals {
				st.Terminals = append(st.Terminals, SessionTerminal(term))
			}

			sw.Tabs = append(sw.Tabs, st)
		}

		session.Windows = append(session.Windows, sw)
	}

	return session, nil
}
