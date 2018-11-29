package word2

import (
	"fmt"
	"testing"
)

var table = []struct {
	input string
	want bool
}{
	{"", true},
	{"a", true},
	{"aa", true},
	{"ab", false},
	{"kayak", true},
	{"detartrated", true},
	{"A man, a plan, a canal: Panama", true},
	{"Evil I did dwell; lewd did I live.", true},
	{"Able was I ere I saw Elba", true},
	{"ёtё", true},
	{"Et se resservir, ivresse reste.", true},
	{"palindrome", false},
	{"desserts", false},
}

func TestIsPalindrome(t *testing.T) {
	for i, test := range table {
		t.Run(fmt.Sprintf("table#%d : IsPalindrome(%q)", i, test.input), func(t *testing.T) {
			if got := IsPalindrome(test.input); got != test.want {
				t.Errorf("IsPalindrome(%q) = %v", test.input, got)
			}
		})
	}
}
