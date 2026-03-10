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
			tabs, err := ghostty.ListTabs(windowID)
			if err != nil {
				return err
			}

			printTabs(tabs)
			return nil
		},
	}

	cmd.Flags().StringVarP(&windowID, "window", "w", "", "window ID")
	cmd.MarkFlagRequired("window")
	cmd.RegisterFlagCompletionFunc("window", completeWindowIDs)
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
			return ghostty.SelectTab(windowID, id)
		},
	}

	cmd.Flags().StringVarP(&windowID, "window", "w", "", "window ID")
	cmd.MarkFlagRequired("window")
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
			return ghostty.CloseTab(windowID, id)
		},
	}

	cmd.Flags().StringVarP(&windowID, "window", "w", "", "window ID")
	cmd.MarkFlagRequired("window")
	cmd.RegisterFlagCompletionFunc("window", completeWindowIDs)
	cmd.Flags().StringVar(&id, "id", "", "tab ID")
	cmd.MarkFlagRequired("id")
	cmd.RegisterFlagCompletionFunc("id", completeTabIDs)
	return cmd
}
