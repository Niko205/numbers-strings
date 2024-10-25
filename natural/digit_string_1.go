package natural

// DigitString1 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Einer-Stelle einer Zahl >= 21 vorkommen würde.
// Außerdem wird bei Ziffern != 0 das Wort "und" angehängt.
//
// Beispiel: Für 3 soll der String "dreiund" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
// Diese Funktion muss nur für den Normalfall (Zahlen >= 21) funktionieren.
func DigitString1(digit int) string {
	s := ""
	switch digit {
	case 0:

	case 1:
		s = "einund"
	case 2:
		s = "zweiund"
	case 3:
		s = "dreiund"
	case 4:
		s = "vierund"
	case 5:
		s = "fünfund"
	case 6:
		s = "sechsund"
	case 7:
		s = "siebenund"
	case 8:
		s = "achtund"
	case 9:
		s = "neunund"
	}
	return s
}
