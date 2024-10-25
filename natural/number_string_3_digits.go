package natural

// NumberString3Digits erwartet eine Zahl 0 <= n <= 999 und liefert den zugehörigen String.
func NumberString3Digits(n int) string {
	s := ""
	if n < 20 {
		s = NumberStringBelow20(n)
	} else {
		s = NumberStringGreater20(n)
	}
	return s
}
