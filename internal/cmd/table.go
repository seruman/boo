package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"code.selman.me/boo/internal/ghostty"
	"github.com/spf13/cobra"
)

func newTabWriter() *tabwriter.Writer {
	return tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
}

func printWindows(windows []ghostty.Window) {
	w := newTabWriter()
	fmt.Fprintln(w, "ID\tNAME\tSELECTED TAB")
	for _, win := range windows {
		fmt.Fprintf(w, "%s\t%s\t%s\n", win.ID, win.Name, win.SelectedTabID)
	}

	w.Flush()
}

func printTabs(tabs []ghostty.Tab) {
	w := newTabWriter()
	fmt.Fprintln(w, "ID\tNAME\tINDEX\tSELECTED")
	for _, t := range tabs {
		fmt.Fprintf(w, "%s\t%s\t%d\t%t\n", t.ID, t.Name, t.Index, t.Selected)
	}

	w.Flush()
}

func printTerminals(terminals []ghostty.Terminal) {
	w := newTabWriter()
	fmt.Fprintln(w, "ID\tNAME\tWORKING DIRECTORY")
	for _, t := range terminals {
		fmt.Fprintf(w, "%s\t%s\t%s\n", t.ID, t.Name, t.WorkingDirectory)
	}

	w.Flush()
}

func printTerminal(t ghostty.Terminal) {
	printTerminals([]ghostty.Terminal{t})
}

func addSurfaceConfigFlags(cmd *cobra.Command, cfg *ghostty.SurfaceConfig) {
	cmd.Flags().StringVarP(&cfg.Command, "command", "c", "", "command to execute")
	cmd.Flags().StringVar(&cfg.WorkingDir, "working-dir", "", "initial working directory")
	cmd.Flags().StringVar(&cfg.InitialInput, "initial-input", "", "input sent to terminal after launch")
	cmd.Flags().Float64Var(&cfg.FontSize, "font-size", 0, "font size in points")
	cmd.Flags().StringSliceVarP(&cfg.EnvVars, "env", "e", nil, "environment variables in KEY=VALUE format (repeatable)")
	cmd.Flags().BoolVar(&cfg.WaitAfterCommand, "wait-after-command", false, "keep terminal open after command exit")
}
