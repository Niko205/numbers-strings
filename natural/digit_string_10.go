package natural

// DigitString10 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Zehner-Stelle einer Zahl >= 21 vorkommen würde.
//
// Beispiel: Für 3 soll der String "dreißig" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
func DigitString10(digit int) string {
	s := ""
	switch digit {
	case 0:

	case 1:

	case 2:
		s = "zwanzig"
	case 3:
		s = "dreißig"
	case 4:
		s = "vierzig"
	case 5:
		s = "fünfzig"
	case 6:
		s = "sechzig"
	case 7:
		s = "siebzig"
	case 8:
		s = "achtzig"
	case 9:
		s = "neunzig"
	}
	return s
}
