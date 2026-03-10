package ghostty

import "testing"

func TestEscapeAppleScript(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{`hello`, `hello`},
		{`say "hi"`, `say \"hi\"`},
		{`path\to\file`, `path\\to\\file`},
		{`"quote\" and \\slash"`, `\"quote\\\" and \\\\slash\"`},
		{``, ``},
		{"line1\nline2", `line1\nline2`},
		{"col1\tcol2", `col1\tcol2`},
		{"cr\rhere", `cr\rhere`},
		{"all\tin\none\r\n", `all\tin\none\r\n`},
	}

	for _, tt := range tests {
		got := escapeAppleScript(tt.in)
		if got != tt.want {
			t.Errorf("escapeAppleScript(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
