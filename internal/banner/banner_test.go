package banner

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	// Create a temporary "banners" directory for testing isolation
	tmpDir := "banners"
	err := os.MkdirAll(tmpDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temp banners directory: %v", err)
	}
	// Clean up the temporary directory after tests finish
	defer os.RemoveAll(tmpDir)

	// Create a mock banner file with mixed Windows (\r\n) and Unix (\n) line endings
	mockContent := "line1\r\nline2\nline3\r\n"
	mockFileName := "test_banner.txt"
	err = os.WriteFile(filepath.Join(tmpDir, mockFileName), []byte(mockContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create mock banner file: %v", err)
	}

	// Define test cases
	tests := []struct {
		name       string
		bannerName string
		wantLines  []string
		wantErr    bool
	}{
		{
			name:       "Successful load and carriage return stripping",
			bannerName: "test_banner.txt",
			wantLines:  []string{"line1", "line2", "line3"},
			wantErr:    false,
		},
		{
			name:       "File does not exist error handling",
			bannerName: "missing_banner.txt",
			wantLines:  nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// FIXED: Changed tt.Load to Load
			gotLines, err := Load(tt.bannerName)

			// Check if error presence matches expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check if the parsed lines match expectation
			if !reflect.DeepEqual(gotLines, tt.wantLines) {
				t.Errorf("Load() gotLines = %v, want %v", gotLines, tt.wantLines)
			}
		})
	}
}
