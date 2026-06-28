package validation

import (
	"testing"
)

func TestValidateText(t *testing.T) {
	tests := []struct {
		name       string
		inputText  string
		wantErr    bool
		errMessage string
	}{
		{
			name:      "Valid standard English text",
			inputText: "Hello World!",
			wantErr:   false,
		},
		{
			name:      "Valid text with safe newlines and carriage returns",
			inputText: "Hello\nWorld\r",
			wantErr:   false,
		},
		{
			name:       "Empty string input",
			inputText:  "",
			wantErr:    true,
			errMessage: "Please enter text to generate ASCII art.",
		},
		{
			name:       "Input with only spaces and tabs",
			inputText:  "   \t  ",
			wantErr:    true,
			errMessage: "Please enter text to generate ASCII art.",
		},
		{
			name:       "Non-ASCII character (Greek character)",
			inputText:  "Γειά",
			wantErr:    true,
			errMessage: "Invalid characters detected. Please use standard English characters only.",
		},
		{
			name:       "Extended ASCII or special symbol outside range 32-126",
			inputText:  "Hello ¥",
			wantErr:    true,
			errMessage: "Invalid characters detected. Please use standard English characters only.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateText(tt.inputText)

			// Check if error presence matches expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Verify the exact error message text matches the UI specification
			if tt.wantErr && err.Error() != tt.errMessage {
				t.Errorf("ValidateText() error message = %q, want %q", err.Error(), tt.errMessage)
			}
		})
	}
}

func TestValidateBanner(t *testing.T) {
	tests := []struct {
		name       string
		bannerName string
		wantErr    bool
		errMessage string
	}{
		{
			name:       "Valid standard banner",
			bannerName: "standard",
			wantErr:    false,
		},
		{
			name:       "Valid shadow banner",
			bannerName: "shadow",
			wantErr:    false,
		},
		{
			name:       "Valid thinkertoy banner",
			bannerName: "thinkertoy",
			wantErr:    false,
		},
		{
			name:       "Invalid banner style injection attempt",
			bannerName: "custom_banner",
			wantErr:    true,
			errMessage: "Please select a valid banner style.",
		},
		{
			name:       "Empty banner string parameter",
			bannerName: "",
			wantErr:    true,
			errMessage: "Please select a valid banner style.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateBanner(tt.bannerName)

			// Check if error presence matches expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateBanner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Verify the error message matches the security specification
			if tt.wantErr && err.Error() != tt.errMessage {
				t.Errorf("ValidateBanner() error message = %q, want %q", err.Error(), tt.errMessage)
			}
		})
	}
}
