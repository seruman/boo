package ghostty

func FocusedTerminal() (Terminal, error) {
	desc, err := execTemplate("app_focused_terminal.applescript", baseParams())
	if err != nil {
		return Terminal{}, err
	}

	return Terminal{
		ID:               desc.DescriptorAtIndex(1).StringValue(),
		Name:             desc.DescriptorAtIndex(2).StringValue(),
		WorkingDirectory: desc.DescriptorAtIndex(3).StringValue(),
	}, nil
}

func FindTerminalsByCwd(query string) ([]Terminal, error) {
	return findTerminals("terminal_find_by_cwd.applescript", query)
}

func FindTerminalsByName(query string) ([]Terminal, error) {
	return findTerminals("terminal_find_by_name.applescript", query)
}

func findTerminals(tmpl, query string) ([]Terminal, error) {
	desc, err := execTemplate(tmpl, struct {
		baseData
		Query string
	}{baseParams(), query})
	if err != nil {
		return nil, err
	}

	count := desc.NumberOfItems()
	terminals := make([]Terminal, 0, count)
	for i := 1; i <= count; i++ {
		item := desc.DescriptorAtIndex(i)
		terminals = append(terminals, Terminal{
			ID:               item.DescriptorAtIndex(1).StringValue(),
			Name:             item.DescriptorAtIndex(2).StringValue(),
			WorkingDirectory: item.DescriptorAtIndex(3).StringValue(),
		})
	}

	return terminals, nil
}

type terminalListParams struct {
	baseData
	WindowID string
	TabID    string
}

func listTerminals(p terminalListParams) ([]Terminal, error) {
	desc, err := execTemplate("terminal_list.applescript", p)
	if err != nil {
		return nil, err
	}

	count := desc.NumberOfItems()
	terminals := make([]Terminal, 0, count)
	for i := 1; i <= count; i++ {
		item := desc.DescriptorAtIndex(i)
		terminals = append(terminals, Terminal{
			ID:               item.DescriptorAtIndex(1).StringValue(),
			Name:             item.DescriptorAtIndex(2).StringValue(),
			WorkingDirectory: item.DescriptorAtIndex(3).StringValue(),
		})
	}

	return terminals, nil
}

func ListTerminals() ([]Terminal, error) {
	return listTerminals(terminalListParams{baseData: baseParams()})
}

func ListTerminalsOfWindow(windowID string) ([]Terminal, error) {
	return listTerminals(terminalListParams{baseData: baseParams(), WindowID: windowID})
}

func ListTerminalsOfTab(windowID, tabID string) ([]Terminal, error) {
	return listTerminals(terminalListParams{baseData: baseParams(), WindowID: windowID, TabID: tabID})
}

func GetTerminal(terminalID string) (Terminal, error) {
	desc, err := execTemplate("terminal_get.applescript", struct {
		baseData
		TerminalID string
	}{baseParams(), terminalID})
	if err != nil {
		return Terminal{}, err
	}

	return Terminal{
		ID:               desc.DescriptorAtIndex(1).StringValue(),
		Name:             desc.DescriptorAtIndex(2).StringValue(),
		WorkingDirectory: desc.DescriptorAtIndex(3).StringValue(),
	}, nil
}

func SplitTerminal(terminalID, direction string, cfg SurfaceConfig) (string, error) {
	if err := validateToken(direction, "direction"); err != nil {
		return "", err
	}

	type params struct {
		surfaceConfigData
		TerminalID string
		Direction  string
	}

	p := params{
		surfaceConfigData: surfaceConfigParams(cfg),
		TerminalID:        terminalID,
		Direction:         direction,
	}

	desc, err := execTemplate("terminal_split.applescript", p)
	if err != nil {
		return "", err
	}

	return desc.StringValue(), nil
}

func FocusTerminal(terminalID string) error {
	_, err := execTemplate("terminal_focus.applescript", struct {
		baseData
		TerminalID string
	}{baseParams(), terminalID})
	return err
}

func CloseTerminal(terminalID string) error {
	_, err := execTemplate("terminal_close.applescript", struct {
		baseData
		TerminalID string
	}{baseParams(), terminalID})
	return err
}

func PerformAction(action, terminalID string) (bool, error) {
	desc, err := execTemplate("terminal_action.applescript", struct {
		baseData
		Action     string
		TerminalID string
	}{baseParams(), action, terminalID})
	if err != nil {
		return false, err
	}

	return desc.BooleanValue(), nil
}
