package ghostty

func ListWindows() ([]Window, error) {
	desc, err := execTemplate("window_list.applescript", baseParams())
	if err != nil {
		return nil, err
	}

	count := desc.NumberOfItems()
	windows := make([]Window, 0, count)
	for i := 1; i <= count; i++ {
		item := desc.DescriptorAtIndex(i)
		windows = append(windows, Window{
			ID:            item.DescriptorAtIndex(1).StringValue(),
			Name:          item.DescriptorAtIndex(2).StringValue(),
			SelectedTabID: item.DescriptorAtIndex(3).StringValue(),
		})
	}

	return windows, nil
}

func NewWindow(cfg SurfaceConfig) (string, error) {
	desc, err := execTemplate("window_new.applescript", surfaceConfigParams(cfg))
	if err != nil {
		return "", err
	}

	return desc.StringValue(), nil
}

func FrontWindow() (Window, error) {
	desc, err := execTemplate("app_front_window.applescript", baseParams())
	if err != nil {
		return Window{}, err
	}

	return Window{
		ID:            desc.DescriptorAtIndex(1).StringValue(),
		Name:          desc.DescriptorAtIndex(2).StringValue(),
		SelectedTabID: desc.DescriptorAtIndex(3).StringValue(),
	}, nil
}

func ActivateWindow(windowID string) error {
	_, err := execTemplate("window_activate.applescript", struct {
		baseData
		WindowID string
	}{baseParams(), windowID})
	return err
}

func CloseWindow(windowID string) error {
	_, err := execTemplate("window_close.applescript", struct {
		baseData
		WindowID string
	}{baseParams(), windowID})
	return err
}
