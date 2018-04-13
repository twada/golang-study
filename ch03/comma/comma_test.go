package comma

import (
	"testing"
)

var table = []struct {
	in  string
	out string
}{
	{"12345", "12,345"},
}

func TestComma(t *testing.T) {
	for _, tt := range table {
		t.Run(tt.in, func(t *testing.T) {
			actual := comma(tt.in)
			if actual != tt.out {
				t.Errorf("got %v\nwant %v", actual, tt.out)
			}
		})
	}
}
