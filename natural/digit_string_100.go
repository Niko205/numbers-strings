package natural

// DigitString100 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Hunderter-Stelle einer Zahl >= 21 vorkommen würde.
//
// Beispiel: Für 3 soll der String "dreihundert" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
func DigitString100(digit int) string {
	s := ""
	switch digit {
	case 0:

	case 1:
		s = "einhundert"
	case 2:
		s = "zweihundert"
	case 3:
		s = "dreihundert"
	case 4:
		s = "vierhundert"
	case 5:
		s = "fünfhundert"
	case 6:
		s = "sechshundert"
	case 7:
		s = "siebenhundert"
	case 8:
		s = "achthundert"
	case 9:
		s = "neunhundert"
	}
	return s
}
