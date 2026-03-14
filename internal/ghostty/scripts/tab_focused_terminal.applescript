tell application id "{{.BundleID}}"
	set t to focused terminal of tab id "{{escape .TabID}}" of window id "{{escape .WindowID}}"
	return {id of t, name of t, working directory of t}
end tell
