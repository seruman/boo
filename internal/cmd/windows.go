package cmd

import (
	"fmt"

	"code.selman.me/boo/internal/ghostty"
	"github.com/spf13/cobra"
)

func newWindowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "window",
		Short: "Manage Ghostty windows",
	}

	cmd.AddCommand(
		newWindowListCmd(),
		newWindowFrontCmd(),
		newWindowNewCmd(),
		newWindowActivateCmd(),
		newWindowCloseCmd(),
	)

	return cmd
}

func newWindowListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all windows",
		RunE: func(cmd *cobra.Command, args []string) error {
			windows, err := ghostty.ListWindows()
			if err != nil {
				return err
			}

			printWindows(windows)
			return nil
		},
	}
}

func newWindowFrontCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "front",
		Short: "Get the frontmost window",
		RunE: func(cmd *cobra.Command, args []string) error {
			w, err := ghostty.FrontWindow()
			if err != nil {
				return err
			}

			printWindows([]ghostty.Window{w})
			return nil
		},
	}
}

func newWindowNewCmd() *cobra.Command {
	var cfg ghostty.SurfaceConfig

	cmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new window",
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := ghostty.NewWindow(cfg)
			if err != nil {
				return err
			}

			fmt.Println(id)
			return nil
		},
	}

	addSurfaceConfigFlags(cmd, &cfg)
	return cmd
}

func newWindowActivateCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "activate",
		Short: "Activate a window",
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedID, err := resolveWindowID(id)
			if err != nil {
				return err
			}

			return ghostty.ActivateWindow(resolvedID)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "window ID (defaults to front window)")
	cmd.RegisterFlagCompletionFunc("id", completeWindowIDs)
	return cmd
}

func newWindowCloseCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "close",
		Short: "Close a window",
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedID, err := resolveWindowID(id)
			if err != nil {
				return err
			}

			return ghostty.CloseWindow(resolvedID)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "window ID (defaults to front window)")
	cmd.RegisterFlagCompletionFunc("id", completeWindowIDs)
	return cmd
}
