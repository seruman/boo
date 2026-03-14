package ghostty

func ListTabs(windowID string) ([]Tab, error) {
	desc, err := execTemplate("tab_list.applescript", struct {
		baseData
		WindowID string
	}{baseParams(), windowID})
	if err != nil {
		return nil, err
	}

	count := desc.NumberOfItems()
	tabs := make([]Tab, 0, count)
	for i := 1; i <= count; i++ {
		item := desc.DescriptorAtIndex(i)
		tabs = append(tabs, Tab{
			ID:       item.DescriptorAtIndex(1).StringValue(),
			Name:     item.DescriptorAtIndex(2).StringValue(),
			Index:    int(item.DescriptorAtIndex(3).Int32Value()),
			Selected: item.DescriptorAtIndex(4).BooleanValue(),
		})
	}

	return tabs, nil
}

func FocusedTerminalOfTab(windowID, tabID string) (Terminal, error) {
	desc, err := execTemplate("tab_focused_terminal.applescript", struct {
		baseData
		WindowID string
		TabID    string
	}{baseParams(), windowID, tabID})
	if err != nil {
		return Terminal{}, err
	}

	return Terminal{
		ID:               desc.DescriptorAtIndex(1).StringValue(),
		Name:             desc.DescriptorAtIndex(2).StringValue(),
		WorkingDirectory: desc.DescriptorAtIndex(3).StringValue(),
	}, nil
}

func NewTab(windowID string, cfg SurfaceConfig) (string, error) {
	type params struct {
		surfaceConfigData
		WindowID string
	}

	p := params{
		surfaceConfigData: surfaceConfigParams(cfg),
		WindowID:          windowID,
	}

	desc, err := execTemplate("tab_new.applescript", p)
	if err != nil {
		return "", err
	}

	return desc.StringValue(), nil
}

func SelectTab(windowID, tabID string) error {
	_, err := execTemplate("tab_select.applescript", struct {
		baseData
		WindowID string
		TabID    string
	}{baseParams(), windowID, tabID})
	return err
}

func CloseTab(windowID, tabID string) error {
	_, err := execTemplate("tab_close.applescript", struct {
		baseData
		WindowID string
		TabID    string
	}{baseParams(), windowID, tabID})
	return err
}
