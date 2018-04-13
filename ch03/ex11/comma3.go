package comma

import (
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	if s[0] == '-' || s[0] == '+' {
		return string(s[0]) + comma(s[1:])
	}
	dot := strings.LastIndex(s, ".")
	if dot != -1 {
		return comma(s[:dot]) + s[dot:]
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
