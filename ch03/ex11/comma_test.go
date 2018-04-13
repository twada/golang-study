package comma

import (
	"testing"
)

var table = []struct {
	in   string
	want string
}{
	{"", ""},
	{"-12", "-12"},
	{"+12", "+12"},
	{"-1.2", "-1.2"},
	{"-12345", "-12,345"},
	{"-123.45", "-123.45"},
	{"-123456", "-123,456"},
	{"-12345.678", "-12,345.678"},
	{"-123456.789", "-123,456.789"},
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
