package cmd

import (
	"fmt"

	"code.selman.me/boo/internal/ghostty"
)

func resolveWindowID(windowID string) (string, error) {
	if windowID != "" {
		return windowID, nil
	}

	window, err := ghostty.FrontWindow()
	if err != nil {
		return "", err
	}

	return window.ID, nil
}

func resolveTabID(windowID, tabID string) (string, string, error) {
	resolvedWindowID, err := resolveWindowID(windowID)
	if err != nil {
		return "", "", err
	}

	if tabID != "" {
		return resolvedWindowID, tabID, nil
	}

	tabs, err := ghostty.ListTabs(resolvedWindowID)
	if err != nil {
		return "", "", err
	}

	for _, tab := range tabs {
		if tab.Selected {
			return resolvedWindowID, tab.ID, nil
		}
	}

	return "", "", fmt.Errorf("no selected tab found for window %s", resolvedWindowID)
}

func resolveTerminalID(terminalID string) (string, error) {
	if terminalID != "" {
		return terminalID, nil
	}

	terminal, err := ghostty.FocusedTerminal()
	if err != nil {
		return "", err
	}

	return terminal.ID, nil
}
