package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"code.selman.me/boo/internal/ghostty"
	"github.com/spf13/cobra"
)

func newSessionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
		Short: "Manage sessions",
	}

	cmd.AddCommand(
		newSessionSaveCmd(),
		newSessionRestoreCmd(),
	)

	return cmd
}

func newSessionSaveCmd() *cobra.Command {
	var output string

	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save current session",
		RunE: func(cmd *cobra.Command, args []string) error {
			session, err := ghostty.CaptureSession()
			if err != nil {
				return err
			}

			data, err := json.MarshalIndent(session, "", "  ")
			if err != nil {
				return fmt.Errorf("marshaling session: %w", err)
			}

			if output == "" {
				_, err = os.Stdout.Write(data)
				fmt.Fprintln(os.Stdout)
				return err
			}

			if err := os.WriteFile(output, data, 0o644); err != nil {
				return err
			}

			fmt.Fprintf(os.Stderr, "session saved to %s\n", output)
			return nil
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "", "write to file instead of stdout")
	return cmd
}

func newSessionRestoreCmd() *cobra.Command {
	var input string
	var dryRun bool

	cmd := &cobra.Command{
		Use:   "restore",
		Short: "Restore a session",
		RunE: func(cmd *cobra.Command, args []string) error {
			var data []byte
			var err error
			if input == "" || input == "-" {
				data, err = io.ReadAll(os.Stdin)
			} else {
				data, err = os.ReadFile(input)
			}

			if err != nil {
				return err
			}

			var session ghostty.Session
			if err := json.Unmarshal(data, &session); err != nil {
				return fmt.Errorf("parsing session: %w", err)
			}

			if dryRun {
				script, err := ghostty.RenderRestoreScript(session)
				if err != nil {
					return err
				}

				fmt.Println(script)
				return nil
			}

			if err := ghostty.RestoreSession(session); err != nil {
				return err
			}

			fmt.Fprintln(os.Stderr, "session restored")
			return nil
		},
	}

	cmd.Flags().StringVarP(&input, "input", "i", "", "input session file (default: stdin)")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "print the generated AppleScript without executing")
	return cmd
}
