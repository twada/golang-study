package basename1

import (
	"testing"
)

var table = []struct {
	in  string
	out string
}{
	{"a", "a"},
	{"a.go", "a"},
	{"a/b/c.go", "c"},
	{"a/b.c.go", "b.c"},
}
func TestBasename(t *testing.T) {
	for _, tt := range table {
		t.Run(tt.in, func(t *testing.T) {
			actual := basename(tt.in)
			if actual != tt.out {
				t.Errorf("got %v\nwant %v", actual, tt.out)
			}
		})
	}
}
