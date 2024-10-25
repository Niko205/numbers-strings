package parse_string

import "strings"

// ParseStringDecimal erwartet einen String, der eine Dezimalzahl repräsentiert, und liefert die zugehörige Zahl.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseStringDecimal(s string) int {
	// TODO
	slist := strings.Split(s, "")
	l := len(slist)
	nrlist := make([]int, l)

	for i := 0; i < len(slist); i++ {
		if slist[i][0] >= '0' && slist[i][0] <= '9' {
			nrlist[i] = int(slist[i][0] - 48)
		} else {
			return -1
		}
	}
	e := 0
	l--
	for i := 0; i < len(nrlist); i++ {
		e = e + nrlist[i]*10 ^ l
		l--
	}
	return e
}
