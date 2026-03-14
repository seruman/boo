package ghostty

import "fmt"

func SetTerminalTitle(title, terminalID string) error {
	return performRequiredAction(titleAction("set_surface_title", title), terminalID)
}

func SetTabTitle(title, terminalID string) error {
	return performRequiredAction(titleAction("set_tab_title", title), terminalID)
}

func titleAction(prefix, title string) string {
	return fmt.Sprintf("%s:%s", prefix, title)
}

func performRequiredAction(action, terminalID string) error {
	ok, err := PerformAction(action, terminalID)
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("ghostty rejected action %q", action)
	}

	return nil
}
