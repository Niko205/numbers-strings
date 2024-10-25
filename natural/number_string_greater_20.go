package natural

// NumberStringGreater20 erwartet eine Zahl >= 20 und liefert den zugehörigen String.
func NumberStringGreater20(n int) string {
	s := ""
	hunderter := 0
	zehner := 0
	einer := 0
	if n < 100 {
		if n < 20 {
			s += NumberStringBelow20(n)
		} else {
			zehner = n / 10
			n = n - (zehner * 10)
			einer = n / 1
			s += DigitString1(einer)
			s += DigitString10(zehner)
		}
	} else {
		if n == 100 {
			hunderter = n / 100
			s += DigitString100(hunderter)
		} else {
			hunderter = n / 100
			n = n - (hunderter * 100)
			s += DigitString100(hunderter)
			if n < 20 {
				s += NumberStringBelow20(n)
			} else {
				zehner = n / 10
				n = n - (zehner * 10)
				einer = n / 1
				s += DigitString1(einer)
				s += DigitString10(zehner)
			}
		}
	}
	return s
}
