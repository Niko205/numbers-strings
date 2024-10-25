package parse_string

// ParseDigit erwartet einen String von 0 bis 9 oder A bis F und liefert die zugehörige Zahl.
// Dabei gilt A=10, B=11, ..., F=15.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseDigit(digit string) int {
	// TODO
	if digit[0] >= '0' && digit[0] <= '9' {
		return int(digit[0] - 48)
	} else if digit[0] >= 'A' && digit[0] <= 'F' {
		return int(digit[0] - 55)
	}
	return -1
}
