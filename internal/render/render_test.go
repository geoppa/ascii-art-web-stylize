package render

import "testing"

func TestGenerate_SpecialCases(t *testing.T) {
	// Φτιάχνουμε ένα μικρό mock banner (9 γραμμές για τον χαρακτήρα 'A')
	// Ο χαρακτήρας 'A' έχει ASCII τιμή 65. Ο τύπος είναι (65-32)*9 + 1 = 298.
	// Για το απλό test μας, γεμίζουμε το slice με 400 γραμμές ώστε
	// να υπάρχουν αρκετές εγγραφές για τους υπολογισμούς των ASCII χαρακτήρων.
	mockBanner := make([]string, 400)
	for i := range mockBanner {
		mockBanner[i] = "###"
	}

	// Έλεγχος ότι το literal "\n" μετατρέπεται σωστά σε πραγματικό newline
	t.Run("Literal Backslash N Conversion", func(t *testing.T) {
		result := Generate("A\\nA", mockBanner)

		if result == Generate("AA", mockBanner) {
			t.Errorf("Generate() failed to convert literal \\n into a real newline")
		}
	})

	// Έλεγχος για την περίπτωση που ο χρήστης στέλνει μόνο Newlines (\n)
	t.Run("All Empty Newlines", func(t *testing.T) {
		input := "\n\n"

		// Το Zone01 περιμένει strings.Repeat("\n", len(parts)-1) -> "\n\n"
		expected := "\n\n"

		result := Generate(input, mockBanner)

		if result != expected {
			t.Errorf("Generate() for empty newlines = %q, want %q", result, expected)
		}
	})

	// Έλεγχος ότι αφαιρείται σωστά το trailing newline στο τέλος του κανονικού κειμένου
	t.Run("No Trailing Newline Check", func(t *testing.T) {
		// Ο χαρακτήρας 'A' (ASCII 65) θα πάρει 8 γραμμές από το mockBanner
		result := Generate("A", mockBanner)

		// Αν το trailing newline αφαιρέθηκε σωστά, ο τελευταίος χαρακτήρας ΔΕΝ πρέπει να είναι '\n'
		if len(result) > 0 && result[len(result)-1] == '\n' {
			t.Errorf("Generate() failed: output has an unwanted trailing newline at the very end")
		}
	})
}