package comma

import (
	"strings"
)

func comma(s string) string {
	var sign, fraction string
	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+") {
		sign = s[0:1]
		s = s[1:]
	}
	dot := strings.LastIndex(s, ".")
	if dot != -1 {
		fraction = s[dot:]
		s = s[:dot]
	}
	return sign + commaRecurse(s) + fraction
}

func commaRecurse(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaRecurse(s[:n-3]) + "," + s[n-3:]
}
