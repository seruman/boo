tell application id "{{.BundleID}}"
	set matches to every terminal whose working directory contains "{{escape .Query}}"
	set output to {}
	repeat with t in matches
		set end of output to {id of t, name of t, working directory of t}
	end repeat
	return output
end tell