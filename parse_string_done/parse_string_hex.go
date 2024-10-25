package parse_string

import (
	"math"
	"strings"
)

// ParseStringDecimal erwartet einen String, der eine Hexadezimalzahl repräsentiert, und liefert die zugehörige Zahl.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseStringHexadecimal(s string) int {
	// TODO
	slist := strings.Split(s, "")
	l := len(slist)
	nrlist := make([]int, l)

	for i := 0; i < len(slist); i++ {
		if slist[i][0] >= '0' && slist[i][0] <= '9' {
			nrlist[i] = int(slist[i][0] - 48)
		} else if slist[i][0] >= 'A' && slist[i][0] <= 'F' {
			nrlist[i] = int(slist[i][0] - 55)
		} else {
			return -1
		}
	}
	e := 0
	l--
	for i := 0; i < len(nrlist); i++ {
		e = e + nrlist[i]*int(math.Pow(16, float64(l)))
		l--
	}
	return e
}
