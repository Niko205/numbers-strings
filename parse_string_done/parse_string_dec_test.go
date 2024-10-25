package parse_string

import "testing"

func TestParseStringDecimal(t *testing.T) {
	if ParseStringDecimal("0") != 0 {
		t.Errorf("ParseStringDecimal(\"0\") == %d, expected 0", ParseStringDecimal("0"))
	}

	if ParseStringDecimal("1") != 1 {
		t.Errorf("ParseStringDecimal(\"1\") == %d, expected 1", ParseStringDecimal("1"))
	}

	if ParseStringDecimal("23") != 23 {
		t.Errorf("ParseStringDecimal(\"23\") == %d, expected 23", ParseStringDecimal("23"))
	}

	if ParseStringDecimal("456") != 456 {
		t.Errorf("ParseStringDecimal(\"456\") == %d, expected 456", ParseStringDecimal("456"))
	}

	if ParseStringDecimal("38") != 38 {
		t.Errorf("ParseStringDecimal(\"38\") == %d, expected 38", ParseStringDecimal("38"))
	}
	if ParseStringDecimal("A") != -1 {
		t.Errorf("ParseStringDecimal(\"A\") == %d, expected -1", ParseStringDecimal("A"))
	}
}
