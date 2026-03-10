tell application id "{{.BundleID}}"
	set output to {}
	repeat with w in every window
		set end of output to {id of w, name of w, id of selected tab of w}
	end repeat
	return output
end tell
