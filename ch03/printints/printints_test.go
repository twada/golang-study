package printints

import (
	"fmt"
	"testing"
)

var table = []struct {
	in  []int
	out string
}{
	{[]int{1,2,3}, "[1, 2, 3]"},
}
func TestComma(t *testing.T) {
	for i, tt := range table {
		t.Run(fmt.Sprintf("table#%d", i), func(t *testing.T) {
			actual := intsToString(tt.in)
			if actual != tt.out {
				t.Errorf("got %v\nwant %v", actual, tt.out)
			}
		})
	}
}
