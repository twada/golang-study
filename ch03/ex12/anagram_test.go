package anagram

import (
	"testing"
)

var table = []struct {
	in1  string
	in2  string
	want bool
}{
	{"canoe", "ocean", true},
	{"canoo", "ocean", false},
}

func TestAnagram(t *testing.T) {
	for _, tt := range table {
		t.Run(tt.in1, func(t *testing.T) {
			got := isAnagram(tt.in1, tt.in2)
			if got != tt.want {
				t.Errorf("got %v\nwant %v", got, tt.want)
			}
		})
	}
}
