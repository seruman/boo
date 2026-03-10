tell application id "{{.BundleID}}"
	set output to {}
	repeat with t in every tab of window id "{{escape .WindowID}}"
		set end of output to {id of t, name of t, index of t, selected of t}
	end repeat
	return output
end tell
