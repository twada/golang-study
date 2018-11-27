package test

import (
	"github.com/twada/golang-study/ch11/charcount"
	"bytes"
	"fmt"
	"testing"
)

var table = []struct {
	input string
	wantCounts map[rune]int
	wantUtflen []int
}{
	{
		"golang",
		map[rune]int {
			'g': 2,
			'o': 1,
			'l': 1,
			'a': 1,
			'n': 1,
		},
		[]int{0, 6, 0, 0, 0},
	},
}

func TestCharcount(t *testing.T) {
	for i, test := range table {
		t.Run(fmt.Sprintf("table#%d : charcount.Count(%q)", i, test.input), func(t *testing.T) {
			in := bytes.NewBufferString(test.input)
			gotCounts, gotUtflen, err := charcount.Count(in)
			if err != nil {
				t.Fatalf("error in charcount.Count: %v", err)
			}
			if len(gotCounts) != len(test.wantCounts) {
				t.Errorf("len(gotCounts) == %d, want %d", len(gotCounts), len(test.wantCounts))
			}
			for k,v := range test.wantCounts {
				got, ok := gotCounts[k]
				if ok {
					if got != v {
						t.Errorf("gotCounts[%q] == %d, want %d", k, got, v)
					}
				} else {
					t.Errorf("gotCounts[%q] does not exist", k)
				}
			}
			if len(gotUtflen) != len(test.wantUtflen) {
				t.Errorf("len(gotUtflen) == %d, want %d", len(gotUtflen), len(test.wantUtflen))
			}
			for i, v := range test.wantUtflen {
				got := gotUtflen[i]
				if got != v {
					t.Errorf("gotUtflen[%d] == %d, want %d", i, got, v)
				}
			}
		})
	}
}
