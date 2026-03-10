tell application id "{{.BundleID}}"
	set w to front window
	return {id of w, name of w, id of selected tab of w}
end tell