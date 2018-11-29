package bench

import (
	"github.com/twada/golang-study/ch11/word2"
	"testing"
)

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		word2.IsPalindrome("A man, a plan, a canal: Panama")
	}
}
