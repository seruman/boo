tell application id "{{.BundleID}}"
	set t to terminal id "{{escape .TerminalID}}"
	return {id of t, name of t, working directory of t}
end tell
