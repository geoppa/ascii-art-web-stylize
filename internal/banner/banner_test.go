package banner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_SuccessAndSanitization(t *testing.T) {
	// 1. Δημιουργούμε έναν προσωρινό φάκελο "banners" για το test
	tmpDir, err := os.MkdirTemp("", "banners_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Αλλάζουμε προσωρινά το working directory στον temp φάκελο,
	// ώστε η Load() να ψάξει στο "./banners/" που δημιουργούμε τώρα
	oldWd, _ := os.Getwd()
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to change working directory: %v", err)
	}
	defer os.Chdir(oldWd) // Επαναφορά του αρχικού directory στο τέλος του test

	err = os.Mkdir("banners", 0755)
	if err != nil {
		t.Fatalf("Failed to create banners subdirectory: %v", err)
	}

	// 2. Δημιουργούμε ένα mock banner αρχείο που περιέχει Windows carriage returns (\r\n)
	// για να προσομοιώσουμε τη συμπεριφορά του thinkertoy.txt
	mockContent := "line1\r\nline2\r\nline3\r\n"
	tmpFile := filepath.Join("banners", "mock_banner.txt")
	err = os.WriteFile(tmpFile, []byte(mockContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write mock banner file: %v", err)
	}

	// 3. Εκτελούμε τη συνάρτηση Load
	lines, err := Load("mock_banner.txt")
	if err != nil {
		t.Fatalf("Load() returned an unexpected error: %v", err)
	}

	// 4. Έλεγχος αν διαβάστηκαν σωστά οι γραμμές (πρέπει να είναι 3)
	if len(lines) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(lines))
	}

	// 5. ΚΡΙΣΙΜΟΣ ΕΛΕΓΧΟΣ: Βεβαιωνόμαστε ότι το '\r' αφαιρέθηκε επιτυχώς
	for _, line := range lines {
		for _, char := range line {
			if char == '\r' {
				t.Errorf("Load() failed to strip out Windows carriage return (\\r)")
			}
		}
	}
}

func TestLoad_FileNotFound(t *testing.T) {
	// Έλεγχος ότι η συνάρτηση επιστρέφει σωστά σφάλμα αν το αρχείο δεν υπάρχει
	_, err := Load("non_existent_banner_file_xyz.txt")
	if err == nil {
		t.Errorf("Expected an error for non-existent file, but got nil")
	}
}
