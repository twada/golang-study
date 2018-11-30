package ex06

import (
	"fmt"
	"math"
	"testing"
)

var table = []struct {
	name string
	sut  func(x uint64) int
}{
	{"original", PopCount},
	{"ex03", PopCountEx03},
	{"ex04", PopCountEx04},
	{"ex05", PopCountEx05},
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
