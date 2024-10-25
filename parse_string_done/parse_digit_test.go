package parse_string

import "testing"

func TestParseDigit(t *testing.T) {
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	expected_digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for i, d := range digits {
		expected := expected_digits[i]
		if ParseDigit(d) != expected {
			t.Errorf("Digit(\"%s\") == %d, expected %d", d, ParseDigit(d), expected)
		}
	}
}
