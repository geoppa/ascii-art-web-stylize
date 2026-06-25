package render

import (
	"strings"
)

// Generate translates input text into visual structural ASCII art blocks
func Generate(text string, banner []string) string {
	var output strings.Builder

	// Turn literal written "\n" text strings into real newline bytes
	text = strings.ReplaceAll(text, "\\n", "\n")

	// Standardize browser carriage returns (\r\n) down to clean standard newlines (\n)
	text = strings.ReplaceAll(text, "\r\n", "\n")

	// Split text on real newline bytes
	parts := strings.Split(text, "\n")

	// Special Audit Case: If the entire input is composed purely of empty newlines,
	// Zone01 expects a matching count of raw newline outputs.
	allEmpty := true
	for _, part := range parts {
		if part != "" {
			allEmpty = false
			break
		}
	}
	if allEmpty && len(parts) > 0 {
		return strings.Repeat("\n", len(parts)-1)
	}

	for i, part := range parts {
		// Handle empty lines between words
		if part == "" {
			// If it's the absolute final line after a trailing newline, skip it
			if i == len(parts)-1 {
				continue
			}
			output.WriteString("\n")
			continue
		}

		// Loop through each of the 8 visual line layers of the characters
		for row := 0; row < 8; row++ {
			for _, char := range part {
				// Calculate the perfect entry baseline index inside the loaded layout array
				start := int(char-32)*9 + 1

				if start+row < len(banner) {
					output.WriteString(banner[start+row])
				}
			}
			output.WriteString("\n")
		}
	}

	return strings.TrimSuffix(output.String(), "\n")
}
