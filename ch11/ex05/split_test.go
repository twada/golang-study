package ex05

import (
	"fmt"
	"strings"
	"testing"
)

var table = []struct {
	s string
	sep string
	want []string
}{
	{"a:b:c", ":", []string{"a","b","c"}},
	{"abc", "", []string{"a","b","c"}},
	{"a b c", "", []string{"a"," ", "b"," ","c"}},
}

func TestSplit(t *testing.T) {
	for i, test := range table {
		t.Run(fmt.Sprintf("table#%d/strings.Split(%q,%q)", i, test.s, test.sep), func(t *testing.T) {
			words := strings.Split(test.s, test.sep)
			if gotlen, wantlen := len(words), len(test.want); gotlen != wantlen {
				t.Errorf("Split(%q, %q) returned %d words, want %d", test.s, test.sep, gotlen, wantlen)
			}
			for i, want := range test.want {
				if got := words[i]; got != want {
					t.Errorf("Split(%q, %q)[%d] = %q, want %q", test.s, test.sep, i, got, want)
				}
			}
		})
	}
}
