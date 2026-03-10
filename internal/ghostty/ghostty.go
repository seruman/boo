package ghostty

type AppInfo struct {
	Name      string
	Version   string
	Frontmost bool
}

type Window struct {
	ID            string
	Name          string
	SelectedTabID string
}

type Tab struct {
	ID       string
	Name     string
	Index    int
	Selected bool
}

type Terminal struct {
	ID               string
	Name             string
	WorkingDirectory string
}

type Session struct {
	Windows []SessionWindow `json:"windows"`
}

type SessionWindow struct {
	ID            string       `json:"id"`
	Name          string       `json:"name"`
	SelectedTabID string       `json:"selected_tab_id"`
	Tabs          []SessionTab `json:"tabs"`
}

type SessionTab struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Index     int               `json:"index"`
	Selected  bool              `json:"selected"`
	Terminals []SessionTerminal `json:"terminals"`
}

type SessionTerminal struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	WorkingDirectory string `json:"working_directory"`
}

type SurfaceConfig struct {
	Command          string
	WorkingDir       string
	InitialInput     string
	FontSize         float64
	EnvVars          []string
	WaitAfterCommand bool
}
