tell application id "{{.BundleID}}" to send mouse position x {{.X}} y {{.Y}}{{if .Modifiers}} modifiers "{{escape .Modifiers}}"{{end}} to terminal id "{{escape .TerminalID}}"
