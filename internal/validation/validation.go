package validation

import (
	"errors"
	"strings"
)

// ValidateText ensures the string is not empty and contains only printable ASCII characters
func ValidateText(text string) error {
	// reject inputs that are only spaces or empty lines
	if strings.TrimSpace(text) == "" {
		return errors.New("Please enter text to generate ASCII art.")
	}

	// loop through every character to protect rendering engine from crashing
	for _, char := range text {
		// allow standard system characters: Newline (\n) and Carriage Return (\r)
		if char == '\n' || char == '\r' {
			continue
		}
		// reject any character outside the standard printable font file index (32 to 126)
		if char < 32 || char > 126 {
			return errors.New("Invalid characters detected. Please use standard English characters only.")
		}
	}

	return nil
}

// ValidateBanner verifies that the chosen layout theme perfectly matches a known system asset
func ValidateBanner(name string) error {
	switch name {
	case "standard", "shadow", "thinkertoy":
		return nil
	default:
		return errors.New("Please select a valid banner style.")
	}
}
