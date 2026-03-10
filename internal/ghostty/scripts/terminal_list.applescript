tell application id "{{.BundleID}}"
	set output to {}
{{- if and .WindowID .TabID}}
	repeat with t in every terminal of tab id "{{escape .TabID}}" of window id "{{escape .WindowID}}"
{{- else if .WindowID}}
	repeat with t in every terminal of window id "{{escape .WindowID}}"
{{- else}}
	repeat with t in every terminal
{{- end}}
		set end of output to {id of t, name of t, working directory of t}
	end repeat
	return output
end tell
