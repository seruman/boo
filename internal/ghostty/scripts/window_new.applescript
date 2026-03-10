tell application id "{{.BundleID}}"
{{- if .HasConfig}}
	set cfg to new surface configuration
{{- if .Command}}
	set command of cfg to "{{escape .Command}}"
{{- end}}
{{- if .WorkingDir}}
	set initial working directory of cfg to "{{escape .WorkingDir}}"
{{- end}}
{{- if .InitialInput}}
	set initial input of cfg to "{{escape .InitialInput}}"
{{- end}}
{{- if .FontSize}}
	set font size of cfg to {{.FontSize}}
{{- end}}
{{- if .WaitAfterCommand}}
	set wait after command of cfg to true
{{- end}}
{{- if .EnvVarsList}}
	set environment variables of cfg to {{"{"}}{{.EnvVarsList}}{{"}"}}
{{- end}}
	set w to new window with configuration cfg
{{- else}}
	set w to new window
{{- end}}
	return id of w
end tell
