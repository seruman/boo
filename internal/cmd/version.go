package cmd

import (
	"fmt"

	"code.selman.me/boo/internal/ghostty"
	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print Ghostty version",
		RunE: func(cmd *cobra.Command, args []string) error {
			v, err := ghostty.GetVersion()
			if err != nil {
				return err
			}

			fmt.Println(v)
			return nil
		},
	}
}

func newInfoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Print Ghostty app info",
		RunE: func(cmd *cobra.Command, args []string) error {
			info, err := ghostty.GetAppInfo()
			if err != nil {
				return err
			}

			w := newTabWriter()
			fmt.Fprintln(w, "NAME\tVERSION\tFRONTMOST")
			fmt.Fprintf(w, "%s\t%s\t%t\n", info.Name, info.Version, info.Frontmost)
			w.Flush()
			return nil
		},
	}
}

func newQuitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "quit",
		Short: "Quit Ghostty",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ghostty.Quit()
		},
	}
}
