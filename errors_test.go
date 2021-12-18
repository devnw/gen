package gen

import "testing"

func Test_ErrIndexMismatch(t *testing.T) {
	err := ErrIndexMismatch[string]{2, "c", "d"}

	expected := "index mismatch: expected c, actual d"
	if err.Error() != expected {
		t.Errorf("Expected error %v, got %v", expected, err)
	}
}

func Test_ErrLengthMismatch(t *testing.T) {
	err := ErrLengthMismatch{3, 2}

	expected := "length mismatch: expected 3, actual 2"
	if err.Error() != expected {
		t.Errorf("Expected error %v, got %v", expected, err)
	}
}
