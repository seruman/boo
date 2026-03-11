package cmd

import (
	"fmt"

	"code.selman.me/boo/internal/ghostty"
	"github.com/spf13/cobra"
)

func completeWindowIDs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	windows, err := ghostty.ListWindows()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	results := make([]string, 0, len(windows))
	for _, w := range windows {
		results = append(results, fmt.Sprintf("%s\t%s", w.ID, w.Name))
	}

	return results, cobra.ShellCompDirectiveNoFileComp
}

func completeTabIDs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	windowID, _ := cmd.Flags().GetString("window")
	if windowID == "" {
		frontWindow, err := ghostty.FrontWindow()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		windowID = frontWindow.ID
	}

	tabs, err := ghostty.ListTabs(windowID)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	results := make([]string, 0, len(tabs))
	for _, t := range tabs {
		results = append(results, fmt.Sprintf("%s\t%s", t.ID, t.Name))
	}

	return results, cobra.ShellCompDirectiveNoFileComp
}

func completeTerminalIDs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	terminals, err := ghostty.ListTerminals()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	results := make([]string, 0, len(terminals))
	for _, t := range terminals {
		results = append(results, fmt.Sprintf("%s\t%s", t.ID, t.Name))
	}

	return results, cobra.ShellCompDirectiveNoFileComp
}
