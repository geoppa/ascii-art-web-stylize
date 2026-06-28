package render

import (
	"strings"
	"testing"
)

// createMockBanner generates a minimal banner structure matching the standard format.
// In the standard font, each character has 8 lines of height + 1 empty line separator = 9 lines total.
// Space (rune 32) is the first character.
func createMockBanner() []string {
	// 1 initial empty line or offset usually found at the start of banner files
	banner := []string{""}

	// We populate mock 9-line blocks for a few characters to use in tests.
	// Character 'A' (rune 65) -> index = (65-32)*9 + 1 = 298
	// Character 'B' (rune 66) -> index = (66-32)*9 + 1 = 307

	// Fill up the array with dummy data up to index 400 so we don't out-of-bounds
	for i := 0; i < 400; i++ {
		banner = append(banner, " . ")
	}

	// Override specific rows for 'A' (lines 298 to 305 are the 8 visual rows)
	for row := 0; row < 8; row++ {
		banner[298+row] = "[A-row" + string(rune('1'+row)) + "]"
	}

	// Override specific rows for 'B' (lines 307 to 314 are the 8 visual rows)
	for row := 0; row < 8; row++ {
		banner[307+row] = "[B-row" + string(rune('1'+row)) + "]"
	}

	return banner
}

func TestGenerate(t *testing.T) {
	banner := createMockBanner()

	tests := []struct {
		name       string
		inputText  string
		wantOutput string
	}{
		{
			name:      "Single character transformation",
			inputText: "A",
			wantOutput: "[A-row1]\n" +
				"[A-row2]\n" +
				"[A-row3]\n" +
				"[A-row4]\n" +
				"[A-row5]\n" +
				"[A-row6]\n" +
				"[A-row7]\n" +
				"[A-row8]",
		},
		{
			name:      "Multiple characters on the same line",
			inputText: "AB",
			wantOutput: "[A-row1][B-row1]\n" +
				"[A-row2][B-row2]\n" +
				"[A-row3][B-row3]\n" +
				"[A-row4][B-row4]\n" +
				"[A-row5][B-row5]\n" +
				"[A-row6][B-row6]\n" +
				"[A-row7][B-row7]\n" +
				"[A-row8][B-row8]",
		},
		{
			name:      "Literal escaped newline resolution",
			inputText: "A\\nB",
			wantOutput: "[A-row1]\n[A-row2]\n[A-row3]\n[A-row4]\n[A-row5]\n[A-row6]\n[A-row7]\n[A-row8]\n" +
				"[B-row1]\n[B-row2]\n[B-row3]\n[B-row4]\n[B-row5]\n[B-row6]\n[B-row7]\n[B-row8]",
		},
		{
			name:      "Windows carriage return stripping",
			inputText: "A\r\nB",
			wantOutput: "[A-row1]\n[A-row2]\n[A-row3]\n[A-row4]\n[A-row5]\n[A-row6]\n[A-row7]\n[A-row8]\n" +
				"[B-row1]\n[B-row2]\n[B-row3]\n[B-row4]\n[B-row5]\n[B-row6]\n[B-row7]\n[B-row8]",
		},
		{
			name:       "Zone01 Audit Case: Pure newlines string",
			inputText:  "\\n\\n",
			wantOutput: "\n\n",
		},
		{
			name:       "Zone01 Audit Case: Actual newlines raw string",
			inputText:  "\n\n",
			wantOutput: "\n\n",
		},
		{
			name:      "Empty block line between valid text",
			inputText: "A\n\nB",
			wantOutput: "[A-row1]\n[A-row2]\n[A-row3]\n[A-row4]\n[A-row5]\n[A-row6]\n[A-row7]\n[A-row8]\n\n" +
				"[B-row1]\n[B-row2]\n[B-row3]\n[B-row4]\n[B-row5]\n[B-row6]\n[B-row7]\n[B-row8]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput := Generate(tt.inputText, banner)
			if gotOutput != tt.wantOutput {
				// Using strings.ReplaceAll in error log to visualize spaces/newlines clearly if mismatch happens
				t.Errorf("Generate() mismatch.\nGot:\n%s\nWant:\n%s",
					strings.ReplaceAll(gotOutput, "\n", "\\n"),
					strings.ReplaceAll(tt.wantOutput, "\n", "\\n"))
			}
		})
	}
}
