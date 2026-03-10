package ghostty

import (
	"strings"
	"testing"
)

func TestGenerateRestoreScript_EmptySession(t *testing.T) {
	session := Session{}
	script, err := generateRestoreScript(session)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(script, "tell application id") {
		t.Fatal("expected applescript wrapper")
	}
}

func TestGenerateRestoreScript_EmptyTab(t *testing.T) {
	session := Session{
		Windows: []SessionWindow{{
			ID:   "w1",
			Name: "win",
			Tabs: []SessionTab{{
				ID:        "t1",
				Name:      "tab",
				Index:     1,
				Terminals: nil,
			}},
		}},
	}

	_, err := generateRestoreScript(session)
	if err != nil {
		t.Fatalf("empty tab should not crash: %v", err)
	}
}

func TestGenerateRestoreScript_AllEmptyTabsWindowSkipped(t *testing.T) {
	session := Session{
		Windows: []SessionWindow{
			{
				ID:   "w1",
				Name: "empty-win",
				Tabs: []SessionTab{
					{ID: "t1", Name: "empty", Index: 1, Terminals: nil},
					{ID: "t2", Name: "also-empty", Index: 2, Terminals: nil},
				},
			},
			{
				ID:   "w2",
				Name: "real-win",
				Tabs: []SessionTab{{
					ID:    "t3",
					Name:  "tab",
					Index: 1,
					Terminals: []SessionTerminal{{
						ID: "a", Name: "s", WorkingDirectory: "/tmp",
					}},
				}},
			},
		},
	}

	script, err := generateRestoreScript(session)
	if err != nil {
		t.Fatal(err)
	}

	if strings.Count(script, "set win to new window") != 1 {
		t.Fatalf("expected exactly 1 window, got:\n%s", script)
	}

	selectIdx := strings.Index(script, "select tab")
	winIdx := strings.Index(script, "set win to new window")
	if selectIdx != -1 && selectIdx < winIdx {
		t.Fatal("select tab appears before set win")
	}
}

func TestGenerateRestoreScript_NewlineInName(t *testing.T) {
	session := Session{
		Windows: []SessionWindow{{
			ID:   "w1",
			Name: "win",
			Tabs: []SessionTab{{
				ID:    "t1",
				Name:  "tab",
				Index: 1,
				Terminals: []SessionTerminal{{
					ID:               "term1",
					Name:             "line1\nline2",
					WorkingDirectory: "/tmp",
				}},
			}},
		}},
	}

	script, err := generateRestoreScript(session)
	if err != nil {
		t.Fatal(err)
	}

	for i, line := range strings.Split(script, "\n") {
		if strings.Contains(line, `input text "`) && !strings.HasSuffix(strings.TrimSpace(line), `" to t1`) {
			t.Fatalf("line %d has broken string literal: %s", i+1, line)
		}
	}

	if !strings.Contains(script, `line1\nline2`) {
		t.Fatalf("expected escaped newline, got:\n%s", script)
	}
}

func TestGenerateRestoreScript_QuotedTerminalName(t *testing.T) {
	session := Session{
		Windows: []SessionWindow{{
			ID:   "w1",
			Name: "win",
			Tabs: []SessionTab{{
				ID:    "t1",
				Name:  "tab",
				Index: 1,
				Terminals: []SessionTerminal{{
					ID:               "term1",
					Name:             `echo "hello"`,
					WorkingDirectory: "/tmp",
				}},
			}},
		}},
	}

	script, err := generateRestoreScript(session)
	if err != nil {
		t.Fatal(err)
	}

	if strings.Contains(script, `was: echo "hello""`) {
		t.Fatal("unescaped quotes in generated script")
	}

	if !strings.Contains(script, `was: echo \"hello\"`) {
		t.Fatalf("expected escaped quotes, got:\n%s", script)
	}
}

func TestGenerateRestoreScript_BackslashInPath(t *testing.T) {
	session := Session{
		Windows: []SessionWindow{{
			ID:   "w1",
			Name: "win",
			Tabs: []SessionTab{{
				ID:    "t1",
				Name:  "tab",
				Index: 1,
				Terminals: []SessionTerminal{{
					ID:               "term1",
					Name:             "shell",
					WorkingDirectory: `/tmp/path with\backslash`,
				}},
			}},
		}},
	}

	script, err := generateRestoreScript(session)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(script, `path with\\backslash`) {
		t.Fatalf("expected escaped backslash, got:\n%s", script)
	}
}

func TestGenerateRestoreScript_BSPLayout(t *testing.T) {
	terminals := make([]SessionTerminal, 4)
	for i := range terminals {
		terminals[i] = SessionTerminal{
			ID:               "id",
			Name:             "term",
			WorkingDirectory: "/tmp",
		}
	}

	session := Session{
		Windows: []SessionWindow{{
			ID:   "w1",
			Name: "win",
			Tabs: []SessionTab{{
				ID:        "t1",
				Name:      "tab",
				Index:     1,
				Terminals: terminals,
			}},
		}},
	}

	script, err := generateRestoreScript(session)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(script, "split t1 direction right") {
		t.Fatal("expected first split to be right")
	}

	if !strings.Contains(script, "split t1 direction down") {
		t.Fatal("expected second split to be down (left half)")
	}

	if !strings.Contains(script, "split t2 direction down") {
		t.Fatal("expected third split to be down (right half)")
	}
}

func TestGenerateRestoreScript_SelectedTab(t *testing.T) {
	session := Session{
		Windows: []SessionWindow{{
			ID:   "w1",
			Name: "win",
			Tabs: []SessionTab{
				{
					ID:    "t1",
					Name:  "tab1",
					Index: 1,
					Terminals: []SessionTerminal{{
						ID: "term1", Name: "s", WorkingDirectory: "/tmp",
					}},
				},
				{
					ID:       "t2",
					Name:     "tab2",
					Index:    2,
					Selected: true,
					Terminals: []SessionTerminal{{
						ID: "term2", Name: "s", WorkingDirectory: "/tmp",
					}},
				},
			},
		}},
	}

	script, err := generateRestoreScript(session)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(script, "select tab (tab 2 of win)") {
		t.Fatalf("expected tab 2 selected, got:\n%s", script)
	}
}

func TestGenerateRestoreScript_MultipleWindows(t *testing.T) {
	session := Session{
		Windows: []SessionWindow{
			{
				ID:   "w1",
				Name: "win1",
				Tabs: []SessionTab{{
					ID:    "t1",
					Name:  "tab1",
					Index: 1,
					Terminals: []SessionTerminal{{
						ID: "a", Name: "s", WorkingDirectory: "/tmp/a",
					}},
				}},
			},
			{
				ID:   "w2",
				Name: "win2",
				Tabs: []SessionTab{{
					ID:    "t2",
					Name:  "tab2",
					Index: 1,
					Terminals: []SessionTerminal{{
						ID: "b", Name: "s", WorkingDirectory: "/tmp/b",
					}},
				}},
			},
		},
	}

	script, err := generateRestoreScript(session)
	if err != nil {
		t.Fatal(err)
	}

	count := strings.Count(script, "set win to new window")
	if count != 2 {
		t.Fatalf("expected 2 new window calls, got %d:\n%s", count, script)
	}

	if strings.Contains(script, "new tab in win") {
		t.Fatal("first tab of each window should create a new window, not a new tab")
	}
}

func TestGenerateRestoreScript_MixedCwds(t *testing.T) {
	session := Session{
		Windows: []SessionWindow{{
			ID:   "w1",
			Name: "win",
			Tabs: []SessionTab{{
				ID:    "t1",
				Name:  "tab",
				Index: 1,
				Terminals: []SessionTerminal{
					{ID: "a", Name: "s", WorkingDirectory: "/home/a"},
					{ID: "b", Name: "s", WorkingDirectory: "/home/b"},
				},
			}},
		}},
	}

	script, err := generateRestoreScript(session)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(script, `"/home/a"`) || !strings.Contains(script, `"/home/b"`) {
		t.Fatalf("expected both cwds preserved, got:\n%s", script)
	}
}
