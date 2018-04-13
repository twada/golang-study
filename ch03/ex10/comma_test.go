package comma

import (
	"testing"
)

var table = []struct {
	in  string
	want string
}{
	{"12345", "12,345"},
	{"123456", "123,456"},
}
func TestComma(t *testing.T) {
	for _, tt := range table {
		t.Run(tt.in, func(t *testing.T) {
			got := comma(tt.in)
			if got != tt.want {
				t.Errorf("got %v\nwant %v", got, tt.want)
			}
		})
	}
}
