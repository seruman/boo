package ghostty

import (
	"fmt"
	"strings"
)

func validateToken(value, name string) error {
	if value == "" {
		return fmt.Errorf("%s must not be empty", name)
	}

	if strings.ContainsAny(value, "\"'\\\n\r\t{}()") {
		return fmt.Errorf("invalid %s: %q", name, value)
	}

	return nil
}
