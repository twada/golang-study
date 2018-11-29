package ex03

import (
	"github.com/twada/golang-study/ch11/word2"
	"math/rand"
	"fmt"
	"time"
	"testing"
	"unicode"
)

func TestRandomNotPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randStr(rng)
		t.Run(fmt.Sprintf("table#%d/IsPalindrome(%q)=true", i, p), func(t *testing.T) {
			if word2.IsPalindrome(p) {
				t.Errorf("IsPalindrome(%q) = true", p)
			}
		})
	}
}

func randStr(rng *rand.Rand) string {
	alphas := upperAlphas()
	n := rng.Intn(23) + 2
	result := make([]rune, n)
	for i := 0; i < n; i++ {
		var r rune
		r, alphas = randPick(alphas, rng)
		if randBool(rng) {
			r = unicode.ToLower(r)
		}
		result[i] = r
	}
	return string(result)
}

func randPick(from []rune, rng *rand.Rand) (rune, []rune) {
	n := rng.Intn(len(from))
	r := from[n]
	to := make([]rune, len(from))
	copy(to, from)
	to = append(to[0:n], to[n+1:]...)
	return r, to
}

func randBool(rng *rand.Rand) bool {
	return rng.Intn(1000) % 2 == 0
}

func upperAlphas() []rune {
	runes := make([]rune, 0)
	for i := 0x0041; i < 0x005b; i++ {
		runes = append(runes, rune(i))
	}
	return runes;
}
