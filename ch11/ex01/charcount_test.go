package test

import (
	"bytes"
	"fmt"
	"github.com/twada/golang-study/ch11/charcount"
	"testing"
)

var table = []struct {
	input      string
	wantCounts map[rune]int
	wantUtflen []int
}{
	{
		"golang",
		map[rune]int{
			'g': 2,
			'o': 1,
			'l': 1,
			'a': 1,
			'n': 1,
		},
		[]int{0, 6, 0, 0, 0},
	},
	{
		"すき家で鯖",
		map[rune]int{
			'す': 1,
			'き': 1,
			'家': 1,
			'で': 1,
			'鯖': 1,
		},
		[]int{0, 0, 0, 5, 0},
	},
	{
		"𠮷野家で𩸽",
		map[rune]int{
			'𠮷': 1,
			'野': 1,
			'家': 1,
			'で': 1,
			'𩸽': 1,
		},
		[]int{0, 0, 0, 3, 2},
	},
}

func TestCharcount(t *testing.T) {
	for i, test := range table {
		t.Run(fmt.Sprintf("table#%d/charcount.Count(%q)", i, test.input), func(t *testing.T) {
			in := bytes.NewBufferString(test.input)
			gotCounts, gotUtflen, err := charcount.Count(in)
			if err != nil {
				t.Fatalf("error in charcount.Count: %v", err)
			}
			if gotlen, wantlen := len(gotCounts), len(test.wantCounts); gotlen != wantlen {
				t.Errorf("len(gotCounts) == %d, want %d", gotlen, wantlen)
			}
			for k, v := range test.wantCounts {
				if got, ok := gotCounts[k]; ok {
					if got != v {
						t.Errorf("gotCounts[%q] == %d, want %d", k, got, v)
					}
				} else {
					t.Errorf("gotCounts[%q] does not exist", k)
				}
			}
			if gotlen, wantlen := len(gotUtflen), len(test.wantUtflen); gotlen != wantlen {
				t.Errorf("len(gotUtflen) == %d, want %d", gotlen, wantlen)
			}
			for i, v := range test.wantUtflen {
				if got := gotUtflen[i]; got != v {
					t.Errorf("gotUtflen[%d] == %d, want %d", i, got, v)
				}
			}
		})
	}
}
