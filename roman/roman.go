package roman

import "strings"

// NToI erwartet eine Zahl und liefert die entsprechende Anzahl an I als String.
func NToI(n int) string {
	return strings.Repeat("I", n)
}

// HINT
// Verwenden Sie die Funktion strings.Repeat.

// NToX erwartet eine Zahl und liefert die entsprechende Anzahl an X als String.
func NToX(n int) string {
	return strings.Repeat("X", n)
}

// HINT
// Verwenden Sie die Funktion strings.Repeat.

// RomanBelow10 erwartet eine Zahl 1 <= n <= 10 und liefert die römische Schreibweise für n als String.
func RomanBelow10(n int) string {
	if n == 0 {
		return ""
	}
	if n <= 3 {
		return NToI(n)
	}
	if n <= 5 {
		return NToI(5-n) + "V"
	}
	if n <= 8 {
		return "V" + NToI(n-5)
	}
	if n <= 10 {
		return NToI(10-n) + "X"
	}
	return ""
}

// HINT
// Unterscheiden Sie die Fälle n == 0, n <= 3, n <= 5, n <= 8 und n <= 10.
// Für jeden dieser Fälle lässt sich mittels der Funktion NToI() ein Teil des Ergebnisses berechnen.

// RomanBelow100 erwartet eine Zahl 1 <= n <= 100 und liefert die römische Schreibweise für n als String.
func RomanBelow100(n int) string {
	if n <= 39 {
		return NToX(n/10) + RomanBelow10(n%10)
	}
	if n <= 59 {
		return NToX((59-n)/10) + "L" + RomanBelow10(n%10)
	}
	if n <= 89 {
		return "L" + NToX((n-50)/10) + RomanBelow10(n%10)
	}
	if n <= 109 {
		return NToX((109-n)/10) + "C" + RomanBelow10(n%10)
	}
	return ""
}

// HINT
// Unterscheiden Sie die Fälle n <= 39, n <= 59, n <= 89 und n <= 109.
// Für jeden dieser Fälle lässt sich mittels der Funktionen NToX() und RomanBelow10()
// ein Teil des Ergebnisses berechnen.
