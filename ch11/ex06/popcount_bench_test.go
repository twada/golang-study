package ex06

import (
	ch02 "github.com/twada/golang-study/ch02/popcount"
	ex03 "github.com/twada/golang-study/ch02/ex03"
	ex04 "github.com/twada/golang-study/ch02/ex04"
	ex05 "github.com/twada/golang-study/ch02/ex05"
	"testing"
)

func BenchmarkPopcountOrig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch02.PopCount(7)
	}
}

func BenchmarkPopcountEx03(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex03.PopCount(7)
	}
}

func BenchmarkPopcountEx04(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex04.PopCount(7)
	}
}

func BenchmarkPopcountEx05(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex05.PopCount(7)
	}
}
