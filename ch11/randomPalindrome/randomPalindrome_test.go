package randp

import (
	"github.com/twada/golang-study/ch11/word2"
	"math/rand"
	"time"
	"testing"
)

func TestRandomPalindromes(t *testing.T) {
	// 疑似乱数生成器を初期化する
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := RandomPalindrome(rng)
		if !word2.IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
