package split

import (
	"fmt"
	"strings"
	"testing"
)

var table = []struct {
	s string
	sep string
	want int
}{
	{"a:b:c", ":", 3},
	{"abc", "", 3},
	{"a b c", "", 5},
}

func TestSplit(t *testing.T) {
	for i, test := range table {
		t.Run(fmt.Sprintf("table#%d : strings.Split(%q,%q)", i, test.s, test.sep), func(t *testing.T) {
			words := strings.Split(test.s, test.sep)
			if got := len(words); got != test.want {
				t.Errorf("Split(%q, %q) returned %d words, want %d", test.s, test.sep, got, test.want)
			}
		})
	}
}
