package comma

import (
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+") {
		return s[0:1] + comma(s[1:])
	}
	dot := strings.LastIndex(s, ".")
	if dot != -1 {
		return comma(s[:dot]) + s[dot:]
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
