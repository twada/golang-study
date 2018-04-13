package anagram

import (
	"fmt"
	"testing"
)

var table = []struct {
	s1  string
	s2  string
	want bool
}{
	{"ocean", "ocean", false},
	{"canoe", "ocean", true},
	{"canoo", "ocean", false},
	{"あとうかい", "かとうあい", true},
	{"あとうかい", "あとうかい", false},
}

func TestAnagram(t *testing.T) {
	for _, tt := range table {
		t.Run(fmt.Sprintf("isAnagram(%v,%v)=>%v", tt.s1, tt.s2, tt.want), func(t *testing.T) {
			got := isAnagram(tt.s1, tt.s2)
			if got != tt.want {
				t.Errorf("got %v\nwant %v", got, tt.want)
			}
		})
	}
}
