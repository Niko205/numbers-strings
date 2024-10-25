package parse_string

import "testing"

func TestParseStringHexadecimal(t *testing.T) {
	if ParseStringHexadecimal("0") != 0 {
		t.Errorf("ParseStringHexadecimal(\"0\") == %d, expected 0", ParseStringHexadecimal("0"))
	}

	if ParseStringHexadecimal("1") != 1 {
		t.Errorf("ParseStringHexadecimal(\"1\") == %d, expected 1", ParseStringHexadecimal("1"))
	}

	if ParseStringHexadecimal("A") != 10 {
		t.Errorf("ParseStringHexadecimal(\"A\") == %d, expected 10", ParseStringHexadecimal("A"))
	}

	if ParseStringHexadecimal("456") != 1110 {
		t.Errorf("ParseStringHexadecimal(\"456\") == %d, expected 1110", ParseStringHexadecimal("456"))
	}

	if ParseStringHexadecimal("38") != 56 {
		t.Errorf("ParseStringHexadecimal(\"38\") == %d, expected 56", ParseStringHexadecimal("38"))
	}
}
