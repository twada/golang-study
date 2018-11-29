package ex06

import (
	"fmt"
	ex03 "github.com/twada/golang-study/ch02/ex03"
	ex04 "github.com/twada/golang-study/ch02/ex04"
	ex05 "github.com/twada/golang-study/ch02/ex05"
	ch02 "github.com/twada/golang-study/ch02/popcount"
	"math"
	"testing"
)

var table = []struct {
	name string
	sut  func(x uint64) int
}{
	{"original", ch02.PopCount},
	{"ex03", ex03.PopCount},
	{"ex04", ex04.PopCount},
	{"ex05", ex05.PopCount},
}

var fixtures = []struct {
	input uint64
	want  int
}{
	{0, 0},
	{math.MaxUint8, 8},
	{math.MaxUint8 + 1, 1},
	{math.MaxUint16, 16},
	{math.MaxUint16 + 1, 1},
	{math.MaxUint32, 32},
	{math.MaxUint32 + 1, 1},
	{math.MaxUint64, 64},
}

func TestPopcount(t *testing.T) {
	for _, test := range fixtures {
		for _, impl := range table {
			t.Run(fmt.Sprintf("%v/PopCount(%d)", impl.name, test.input), func(t *testing.T) {
				if got := impl.sut(test.input); got != test.want {
					t.Errorf("%v/PopCount(%d) = %d, want %d", impl.name, test.input, got, test.want)
				}
			})
		}
	}
}

func BenchmarkPopcount(b *testing.B) {
	for _, impl := range table {
		for _, test := range fixtures {
			b.Run(fmt.Sprintf("%v/PopCount(%d)", impl.name, test.input), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					impl.sut(test.input)
				}
			})
		}
	}
}
