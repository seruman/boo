tell application id "{{.BundleID}}"
	set t to focused terminal of selected tab of front window
	return {id of t, name of t, working directory of t}
end tell