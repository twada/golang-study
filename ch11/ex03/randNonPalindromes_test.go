package ex03

import (
	"github.com/twada/golang-study/ch11/word2"
	"math/rand"
	"time"
	"testing"
)

func randStr(rng *rand.Rand) string {
	n := rng.Intn(20) + 4
	runes := make([]rune, n)
	runes = append(runes, rune('G'))
	runes = append(runes, rune('o'))
	for i := 2; i < n; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
	}
	runes = append(runes, rune('G'))
	runes = append(runes, rune('o'))
	return string(runes)
}

func TestRandomNotPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randStr(rng)
		if word2.IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}
