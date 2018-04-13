package comma

import (
	"bytes"
)

func comma(s string) string {
	var buf bytes.Buffer
	for i, j := len(s)-1, 1; i >= 0; i, j = i-1, j+1 {
		buf.WriteByte(s[i])
		if j%3 == 0 && i != 0 {
			buf.WriteString(",")
		}
	}
	return reversed(buf.String())
}

func reversed(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
