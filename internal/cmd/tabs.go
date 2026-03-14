package cmd

import (
	"fmt"

	"code.selman.me/boo/internal/ghostty"
	"github.com/spf13/cobra"
)

func newTabCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tab",
		Short: "Manage Ghostty tabs",
	}

	cmd.AddCommand(
		newTabListCmd(),
		newTabFocusedTerminalCmd(),
		newTabNewCmd(),
		newTabSelectCmd(),
		newTabCloseCmd(),
	)

	return cmd
}

func newTabListCmd() *cobra.Command {
	var windowID string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List tabs in a window",
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedWindowID, err := resolveWindowID(windowID)
			if err != nil {
				return err
			}

			tabs, err := ghostty.ListTabs(resolvedWindowID)
			if err != nil {
				return err
			}

			printTabs(tabs)
			return nil
		},
	}

	cmd.Flags().StringVarP(&windowID, "window", "w", "", "window ID (defaults to front window)")
	cmd.RegisterFlagCompletionFunc("window", completeWindowIDs)
	return cmd
}

func newTabFocusedTerminalCmd() *cobra.Command {
	var windowID, id string

	cmd := &cobra.Command{
		Use:   "focused-terminal",
		Short: "Get the focused terminal of a tab",
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedWindowID, resolvedTabID, err := resolveTabID(windowID, id)
			if err != nil {
				return err
			}

			terminal, err := ghostty.FocusedTerminalOfTab(resolvedWindowID, resolvedTabID)
			if err != nil {
				return err
			}

			printTerminal(terminal)
			return nil
		},
	}

	cmd.Flags().StringVarP(&windowID, "window", "w", "", "window ID (defaults to front window)")
	cmd.RegisterFlagCompletionFunc("window", completeWindowIDs)
	cmd.Flags().StringVar(&id, "id", "", "tab ID (defaults to selected tab in window)")
	cmd.RegisterFlagCompletionFunc("id", completeTabIDs)
	return cmd
}

func newTabNewCmd() *cobra.Command {
	var windowID string
	var cfg ghostty.SurfaceConfig

	cmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new tab",
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := ghostty.NewTab(windowID, cfg)
			if err != nil {
				return err
			}

			fmt.Println(id)
			return nil
		},
	}

	cmd.Flags().StringVarP(&windowID, "window", "w", "", "window ID (optional)")
	cmd.RegisterFlagCompletionFunc("window", completeWindowIDs)
	addSurfaceConfigFlags(cmd, &cfg)
	return cmd
}

func newTabSelectCmd() *cobra.Command {
	var windowID, id string

	cmd := &cobra.Command{
		Use:   "select",
		Short: "Select a tab",
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedWindowID, err := resolveWindowID(windowID)
			if err != nil {
				return err
			}

			if id == "" {
				return fmt.Errorf("tab ID is required")
			}

			return ghostty.SelectTab(resolvedWindowID, id)
		},
	}

	cmd.Flags().StringVarP(&windowID, "window", "w", "", "window ID (defaults to front window)")
	cmd.RegisterFlagCompletionFunc("window", completeWindowIDs)
	cmd.Flags().StringVar(&id, "id", "", "tab ID")
	cmd.MarkFlagRequired("id")
	cmd.RegisterFlagCompletionFunc("id", completeTabIDs)
	return cmd
}

func newTabCloseCmd() *cobra.Command {
	var windowID, id string

	cmd := &cobra.Command{
		Use:   "close",
		Short: "Close a tab",
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedWindowID, resolvedTabID, err := resolveTabID(windowID, id)
			if err != nil {
				return err
			}

			return ghostty.CloseTab(resolvedWindowID, resolvedTabID)
		},
	}

	cmd.Flags().StringVarP(&windowID, "window", "w", "", "window ID (defaults to front window)")
	cmd.RegisterFlagCompletionFunc("window", completeWindowIDs)
	cmd.Flags().StringVar(&id, "id", "", "tab ID (defaults to selected tab in window)")
	cmd.RegisterFlagCompletionFunc("id", completeTabIDs)
	return cmd
}
