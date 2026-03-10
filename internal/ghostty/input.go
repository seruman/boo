package ghostty

func InputText(text, terminalID string) error {
	_, err := execTemplate("terminal_input.applescript", struct {
		baseData
		Text       string
		TerminalID string
	}{baseParams(), text, terminalID})
	return err
}

func SendKey(key, terminalID, action, modifiers string) error {
	if action != "" {
		if err := validateToken(action, "action"); err != nil {
			return err
		}
	}

	_, err := execTemplate("terminal_send_key.applescript", struct {
		baseData
		Key        string
		TerminalID string
		Action     string
		Modifiers  string
	}{baseParams(), key, terminalID, action, modifiers})
	return err
}

func SendMouseButton(button, terminalID, action, modifiers string) error {
	if err := validateToken(button, "button"); err != nil {
		return err
	}

	if action != "" {
		if err := validateToken(action, "action"); err != nil {
			return err
		}
	}

	_, err := execTemplate("terminal_send_mouse_button.applescript", struct {
		baseData
		Button     string
		TerminalID string
		Action     string
		Modifiers  string
	}{baseParams(), button, terminalID, action, modifiers})
	return err
}

func SendMousePosition(x, y float64, terminalID, modifiers string) error {
	_, err := execTemplate("terminal_send_mouse_pos.applescript", struct {
		baseData
		X          float64
		Y          float64
		TerminalID string
		Modifiers  string
	}{baseParams(), x, y, terminalID, modifiers})
	return err
}

func SendMouseScroll(x, y float64, terminalID string, precision bool, momentum string) error {
	if momentum != "" {
		if err := validateToken(momentum, "momentum"); err != nil {
			return err
		}
	}

	_, err := execTemplate("terminal_send_mouse_scroll.applescript", struct {
		baseData
		X          float64
		Y          float64
		TerminalID string
		Precision  bool
		Momentum   string
	}{baseParams(), x, y, terminalID, precision, momentum})
	return err
}
