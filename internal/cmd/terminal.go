package cmd

import (
	"fmt"

	"code.selman.me/boo/internal/ghostty"
	"github.com/spf13/cobra"
)

func newTerminalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "terminal",
		Short: "Manage Ghostty terminals",
	}

	cmd.AddCommand(
		newTerminalListCmd(),
		newTerminalGetCmd(),
		newTerminalSetTitleCmd(),
		newTerminalFindCmd(),
		newTerminalSplitCmd(),
		newTerminalFocusCmd(),
		newTerminalCloseCmd(),
		newTerminalActionCmd(),
		newTerminalInputCmd(),
		newTerminalSendKeyCmd(),
		newTerminalSendMouseButtonCmd(),
		newTerminalSendMousePosCmd(),
		newTerminalSendMouseScrollCmd(),
	)

	return cmd
}

func newTerminalListCmd() *cobra.Command {
	var windowID, tabID string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List terminals",
		RunE: func(cmd *cobra.Command, args []string) error {
			var terminals []ghostty.Terminal
			var err error
			switch {
			case tabID != "":
				resolvedWindowID, resolveErr := resolveWindowID(windowID)
				if resolveErr != nil {
					return resolveErr
				}
				terminals, err = ghostty.ListTerminalsOfTab(resolvedWindowID, tabID)
			case windowID != "":
				terminals, err = ghostty.ListTerminalsOfWindow(windowID)
			default:
				terminals, err = ghostty.ListTerminals()
			}

			if err != nil {
				return err
			}

			printTerminals(terminals)
			return nil
		},
	}

	cmd.Flags().StringVarP(&windowID, "window", "w", "", "scope to window ID")
	cmd.RegisterFlagCompletionFunc("window", completeWindowIDs)
	cmd.Flags().StringVarP(&tabID, "tab", "t", "", "scope to tab ID (uses front window if --window is omitted)")
	cmd.RegisterFlagCompletionFunc("tab", completeTabIDs)
	return cmd
}

func newTerminalGetCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a terminal",
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedID, err := resolveTerminalID(id)
			if err != nil {
				return err
			}

			t, err := ghostty.GetTerminal(resolvedID)
			if err != nil {
				return err
			}

			printTerminal(t)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "terminal ID (defaults to focused terminal)")
	cmd.RegisterFlagCompletionFunc("id", completeTerminalIDs)
	return cmd
}

func newTerminalSetTitleCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "set-title [title]",
		Short: "Set or clear a terminal title",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedID, err := resolveTerminalID(id)
			if err != nil {
				return err
			}

			title := ""
			if len(args) > 0 {
				title = args[0]
			}

			return ghostty.SetTerminalTitle(title, resolvedID)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "terminal ID (defaults to focused terminal)")
	cmd.RegisterFlagCompletionFunc("id", completeTerminalIDs)
	return cmd
}

func newTerminalFindCmd() *cobra.Command {
	var byCwd, byName string

	cmd := &cobra.Command{
		Use:   "find",
		Short: "Find terminals",
		RunE: func(cmd *cobra.Command, args []string) error {
			var terminals []ghostty.Terminal
			var err error
			switch {
			case byCwd != "":
				terminals, err = ghostty.FindTerminalsByCwd(byCwd)
			case byName != "":
				terminals, err = ghostty.FindTerminalsByName(byName)
			default:
				return fmt.Errorf("specify --cwd or --name")
			}

			if err != nil {
				return err
			}

			printTerminals(terminals)
			return nil
		},
	}

	cmd.Flags().StringVar(&byCwd, "cwd", "", "search by working directory (substring match)")
	cmd.Flags().StringVar(&byName, "name", "", "search by terminal name (substring match)")
	return cmd
}

func newTerminalSplitCmd() *cobra.Command {
	var id, direction string
	var cfg ghostty.SurfaceConfig

	cmd := &cobra.Command{
		Use:   "split",
		Short: "Split a terminal",
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedID, err := resolveTerminalID(id)
			if err != nil {
				return err
			}

			newID, err := ghostty.SplitTerminal(resolvedID, direction, cfg)
			if err != nil {
				return err
			}

			fmt.Println(newID)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "terminal ID (defaults to focused terminal)")
	cmd.RegisterFlagCompletionFunc("id", completeTerminalIDs)
	cmd.Flags().StringVarP(&direction, "direction", "d", "right", "split direction: right, left, down, up")
	addSurfaceConfigFlags(cmd, &cfg)
	return cmd
}

func newTerminalFocusCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "focus",
		Short: "Focus a terminal",
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedID, err := resolveTerminalID(id)
			if err != nil {
				return err
			}

			return ghostty.FocusTerminal(resolvedID)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "terminal ID (defaults to focused terminal)")
	cmd.RegisterFlagCompletionFunc("id", completeTerminalIDs)
	return cmd
}

func newTerminalCloseCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "close",
		Short: "Close a terminal",
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedID, err := resolveTerminalID(id)
			if err != nil {
				return err
			}

			return ghostty.CloseTerminal(resolvedID)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "terminal ID (defaults to focused terminal)")
	cmd.RegisterFlagCompletionFunc("id", completeTerminalIDs)
	return cmd
}

func newTerminalActionCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "action [action-string]",
		Short: "Perform a keybind action",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedID, err := resolveTerminalID(id)
			if err != nil {
				return err
			}

			ok, err := ghostty.PerformAction(args[0], resolvedID)
			if err != nil {
				return err
			}

			fmt.Println(ok)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "terminal ID (defaults to focused terminal)")
	cmd.RegisterFlagCompletionFunc("id", completeTerminalIDs)
	return cmd
}

func newTerminalInputCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "input [text]",
		Short: "Paste text to a terminal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedID, err := resolveTerminalID(id)
			if err != nil {
				return err
			}

			return ghostty.InputText(args[0], resolvedID)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "terminal ID (defaults to focused terminal)")
	cmd.RegisterFlagCompletionFunc("id", completeTerminalIDs)
	return cmd
}

func newTerminalSendKeyCmd() *cobra.Command {
	var id, action, modifiers string

	cmd := &cobra.Command{
		Use:   "send-key [key]",
		Short: "Send a key event",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedID, err := resolveTerminalID(id)
			if err != nil {
				return err
			}

			return ghostty.SendKey(args[0], resolvedID, action, modifiers)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "terminal ID (defaults to focused terminal)")
	cmd.RegisterFlagCompletionFunc("id", completeTerminalIDs)
	cmd.Flags().StringVarP(&action, "action", "a", "", "press or release (default: press)")
	cmd.Flags().StringVarP(&modifiers, "modifiers", "m", "", "comma-separated modifiers: shift,control,option,command")
	return cmd
}

func newTerminalSendMouseButtonCmd() *cobra.Command {
	var id, button, action, modifiers string

	cmd := &cobra.Command{
		Use:   "send-mouse-button",
		Short: "Send a mouse button event",
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedID, err := resolveTerminalID(id)
			if err != nil {
				return err
			}

			return ghostty.SendMouseButton(button, resolvedID, action, modifiers)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "terminal ID (defaults to focused terminal)")
	cmd.RegisterFlagCompletionFunc("id", completeTerminalIDs)
	cmd.Flags().StringVarP(&button, "button", "b", "left button", "mouse button: left button, right button, middle button")
	cmd.Flags().StringVarP(&action, "action", "a", "", "press or release (default: press)")
	cmd.Flags().StringVarP(&modifiers, "modifiers", "m", "", "comma-separated modifiers: shift,control,option,command")
	return cmd
}

func newTerminalSendMousePosCmd() *cobra.Command {
	var id, modifiers string
	var x, y float64

	cmd := &cobra.Command{
		Use:   "send-mouse-pos",
		Short: "Send a mouse position event",
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedID, err := resolveTerminalID(id)
			if err != nil {
				return err
			}

			return ghostty.SendMousePosition(x, y, resolvedID, modifiers)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "terminal ID (defaults to focused terminal)")
	cmd.RegisterFlagCompletionFunc("id", completeTerminalIDs)
	cmd.Flags().Float64Var(&x, "x", 0, "horizontal position in pixels")
	cmd.MarkFlagRequired("x")
	cmd.Flags().Float64Var(&y, "y", 0, "vertical position in pixels")
	cmd.MarkFlagRequired("y")
	cmd.Flags().StringVarP(&modifiers, "modifiers", "m", "", "comma-separated modifiers: shift,control,option,command")
	return cmd
}

func newTerminalSendMouseScrollCmd() *cobra.Command {
	var id, momentum string
	var x, y float64
	var precision bool

	cmd := &cobra.Command{
		Use:   "send-mouse-scroll",
		Short: "Send a mouse scroll event",
		RunE: func(cmd *cobra.Command, args []string) error {
			resolvedID, err := resolveTerminalID(id)
			if err != nil {
				return err
			}

			return ghostty.SendMouseScroll(x, y, resolvedID, precision, momentum)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "terminal ID (defaults to focused terminal)")
	cmd.RegisterFlagCompletionFunc("id", completeTerminalIDs)
	cmd.Flags().Float64Var(&x, "dx", 0, "horizontal scroll delta")
	cmd.Flags().Float64Var(&y, "dy", 0, "vertical scroll delta")
	cmd.Flags().BoolVar(&precision, "precision", false, "high-precision scroll (e.g. trackpad)")
	cmd.Flags().StringVar(&momentum, "momentum", "", "momentum phase: none, began, changed, ended, cancelled, may begin, stationary")
	return cmd
}
