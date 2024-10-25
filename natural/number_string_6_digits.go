package natural

// NumberString6Digits erwartet eine Zahl 0 <= n <= 999999 und liefert den zugehörigen String.
func NumberString6Digits(n int) string {
	s := ""
	tausender := 0
	if n < 1000 {
		s = NumberString3Digits(n)
	} else {
		tausender = n / 1000
		s += NumberString3Digits(n)
		s += "tausend"
		n = n - (tausender * 1000)
		s += NumberString3Digits(n)
	}
	return s
}
