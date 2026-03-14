package ghostty

import "testing"

func TestTitleActionFormatting(t *testing.T) {
	tests := []struct {
		name   string
		prefix string
		title  string
		want   string
	}{
		{name: "terminal", prefix: "set_surface_title", title: "logs", want: "set_surface_title:logs"},
		{name: "tab", prefix: "set_tab_title", title: "prod:blue", want: "set_tab_title:prod:blue"},
		{name: "clear terminal", prefix: "set_surface_title", title: "", want: "set_surface_title:"},
		{name: "clear tab", prefix: "set_tab_title", title: "", want: "set_tab_title:"},
	}

	for _, tt := range tests {
		got := titleAction(tt.prefix, tt.title)
		if got != tt.want {
			t.Fatalf("%s: got %q want %q", tt.name, got, tt.want)
		}
	}
}
