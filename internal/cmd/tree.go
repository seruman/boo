package cmd

import (
	"fmt"
	"os"

	"code.selman.me/boo/internal/ghostty"
	"github.com/spf13/cobra"
)

func newTreeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tree",
		Short: "Show hierarchy",
		RunE: func(cmd *cobra.Command, args []string) error {
			windows, err := ghostty.ListWindows()
			if err != nil {
				return err
			}

			for wi, win := range windows {
				lastWin := wi == len(windows)-1
				winPrefix := "├── "
				winChildPrefix := "│   "
				if lastWin {
					winPrefix = "└── "
					winChildPrefix = "    "
				}

				fmt.Fprintf(os.Stdout, "%s%s (%s)\n", winPrefix, win.Name, win.ID)

				tabs, err := ghostty.ListTabs(win.ID)
				if err != nil {
					fmt.Fprintf(os.Stdout, "%s└── (error listing tabs)\n", winChildPrefix)
					continue
				}

				for ti, tab := range tabs {
					lastTab := ti == len(tabs)-1
					tabPrefix := winChildPrefix + "├── "
					tabChildPrefix := winChildPrefix + "│   "
					if lastTab {
						tabPrefix = winChildPrefix + "└── "
						tabChildPrefix = winChildPrefix + "    "
					}

					sel := ""
					if tab.Selected {
						sel = " *"
					}

					fmt.Fprintf(os.Stdout, "%s%s [%d]%s (%s)\n", tabPrefix, tab.Name, tab.Index, sel, tab.ID)

					terminals, err := ghostty.ListTerminalsOfTab(win.ID, tab.ID)
					if err != nil {
						fmt.Fprintf(os.Stdout, "%s└── (error listing terminals)\n", tabChildPrefix)
						continue
					}

					for ti, term := range terminals {
						lastTerm := ti == len(terminals)-1
						termPrefix := tabChildPrefix + "├── "
						if lastTerm {
							termPrefix = tabChildPrefix + "└── "
						}

						fmt.Fprintf(os.Stdout, "%s%s  %s (%s)\n", termPrefix, term.Name, term.WorkingDirectory, term.ID)
					}
				}
			}

			return nil
		},
	}
}
