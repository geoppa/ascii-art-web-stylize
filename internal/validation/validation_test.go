package validation

import "testing"

func TestValidateText(t *testing.T) {
	// Ορίζουμε test cases με τα αναμενόμενα αποτελέσματα
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"Valid English Text", "Hello World 123!", false},
		{"Valid with Newlines", "Hello\nWorld", false},
		{"Invalid Greek Characters", "Καλημέρα", true},
		{"Invalid Emoji", "Hello 👋", true},
		{"Invalid Only Newlines", "\n\n", true},
		{"Invalid Empty Text", "", true},
		{"Invalid Only Spaces", "   ", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateText(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateBanner(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"Valid Standard", "standard", false},
		{"Valid Shadow", "shadow", false},
		{"Valid Thinkertoy", "thinkertoy", false},
		{"Invalid Banner Name", "random", true},
		{"Invalid Directory Traversal Attack", "../../etc/passwd", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateBanner(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateBanner() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
