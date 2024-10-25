package natural

// NumberStringBelow20 erwartet eine Zahl >= 20 und liefert den zugehörigen String.
func NumberStringBelow20(n int) string {
	s := ""
	switch n {
	case 0:
		s = "null"
	case 1:
		s = "eins"
	case 2:
		s = "zwei"
	case 3:
		s = "drei"
	case 4:
		s = "vier"
	case 5:
		s = "fünf"
	case 6:
		s = "sechs"
	case 7:
		s = "sieben"
	case 8:
		s = "acht"
	case 9:
		s = "neun"
	case 10:
		s = "zehn"
	case 11:
		s = "elf"
	case 12:
		s = "zwölf"
	case 13:
		s = "dreizehn"
	case 14:
		s = "vierzehn"
	case 15:
		s = "fünfzehn"
	case 16:
		s = "sechzehn"
	case 17:
		s = "siebzehn"
	case 18:
		s = "achtzehn"
	case 19:
		s = "neunzehn"
	}
	return s
}
