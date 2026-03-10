tell application id "{{.BundleID}}" to send mouse scroll x {{.X}} y {{.Y}}{{if .Precision}} precision true{{end}}{{if .Momentum}} momentum {{.Momentum}}{{end}} to terminal id "{{escape .TerminalID}}"
