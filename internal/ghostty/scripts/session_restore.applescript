tell application id "{{.BundleID}}"
	activate
{{range $wi, $win := .Windows}}
{{- range $ti, $tab := $win.Tabs}}

	-- Tab {{$tab.Index}}: {{$tab.Name}} ({{len $tab.Terminals}} terminal(s))
{{- range $tab.Terminals}}
	--   [hint] {{.Name}}  ({{.WorkingDirectory}})
{{- end}}
{{- $first := index $tab.Terminals 0}}
	set cfg0 to new surface configuration
	set initial working directory of cfg0 to "{{escape $first.WorkingDirectory}}"
{{- if eq $ti 0}}
	set win to new window with configuration cfg0
	set t1 to terminal 1 of selected tab of win
{{- else}}
	set currentTab to new tab in win with configuration cfg0
	set t1 to terminal 1 of currentTab
{{- end}}
{{- range $tab.Splits}}
	set {{.CfgVar}} to new surface configuration
	set initial working directory of {{.CfgVar}} to "{{escape .WorkingDirectory}}"
	set {{.NewVar}} to split {{.ParentVar}} direction {{.Direction}} with configuration {{.CfgVar}}
{{- end}}
{{- range $tab.Hints}}
	input text "{{escape .Text}}" to {{.Var}}
{{- end}}
{{- end}}
{{- if gt $win.SelectedTabIndex 0}}

	select tab (tab {{$win.SelectedTabIndex}} of win)
{{- end}}
{{- end}}
end tell
