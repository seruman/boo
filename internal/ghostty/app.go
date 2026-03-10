package ghostty

func GetAppInfo() (AppInfo, error) {
	desc, err := execTemplate("app_info.applescript", baseParams())
	if err != nil {
		return AppInfo{}, err
	}

	return AppInfo{
		Name:      desc.DescriptorAtIndex(1).StringValue(),
		Version:   desc.DescriptorAtIndex(2).StringValue(),
		Frontmost: desc.DescriptorAtIndex(3).BooleanValue(),
	}, nil
}

func GetVersion() (string, error) {
	desc, err := execTemplate("app_version.applescript", baseParams())
	if err != nil {
		return "", err
	}

	return desc.StringValue(), nil
}

func Quit() error {
	_, err := execTemplate("app_quit.applescript", baseParams())
	return err
}
