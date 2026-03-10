package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "boo",
		Short: "Control Ghostty via AppleScript",
	}

	rootCmd.AddCommand(
		newTreeCmd(),
		newSessionCmd(),
		newWindowCmd(),
		newTabCmd(),
		newTerminalCmd(),
		newInfoCmd(),
		newVersionCmd(),
		newQuitCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
