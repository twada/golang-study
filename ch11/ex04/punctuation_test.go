package ex04

import (
	word "github.com/twada/golang-study/ch11/word2"
	"math/rand"
	"fmt"
	"time"
	"testing"
	"unicode"
)

func TestIsPalindromeTreatsPunctuations(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randNonPalindrome(rng)
		t.Run(fmt.Sprintf("table#%d/IsPalindrome(%q)=true", i, p), func(t *testing.T) {
			if word.IsPalindrome(p) {
				t.Errorf("IsPalindrome(%q) = true", p)
			}
		})
	}
}

func TestNonPalindromeTreatsPunctuations(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randPalindrome(rng)
		t.Run(fmt.Sprintf("table#%d/IsPalindrome(%q)=false", i, p), func(t *testing.T) {
			if !word.IsPalindrome(p) {
				t.Errorf("IsPalindrome(%q) = false", p)
			}
		})
	}
}

func randPalindrome(rng *rand.Rand) string {
	alphas := upperAlphas()
	puncs := punctuations()
	n := rng.Intn(25) // 24 までのランダムな長さ
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		var r rune
		if randBool(rng) {
			r = randChoose(alphas, rng)
			if randBool(rng) {
				r = unicode.ToLower(r)
			}
		} else {
			r = randChoose(puncs, rng)
		}
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func randNonPalindrome(rng *rand.Rand) string {
	alphas := upperAlphas()
	puncs := punctuations()
	// n := rng.Intn(21) + 4
	n := 26
	result := make([]rune, n)
	for i := 0; i < n; i++ {
		var r rune
		if randBool(rng) {
			r, alphas = randRemove(alphas, rng)
			if randBool(rng) {
				r = unicode.ToLower(r)
			}
		} else {
			r, puncs = randRemove(puncs, rng)
		}
		result[i] = r
	}
	return string(result)
}

func randRemove(from []rune, rng *rand.Rand) (rune, []rune) {
	n := rng.Intn(len(from))
	r := from[n]
	to := append(from[0:n], from[n+1:]...)
	return r, to
}

func randChoose(from []rune, rng *rand.Rand) rune {
	n := rng.Intn(len(from))
	return from[n]
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

func punctuations() []rune {
	runes := make([]rune, 0)
	for i := 0x0020; i < 0x0030; i++ {
		runes = append(runes, rune(i))
	}
	for i := 0x005b; i < 0x0061; i++ {
		runes = append(runes, rune(i))
	}
	for i := 0x007b; i < 0x007f; i++ {
		runes = append(runes, rune(i))
	}
	return runes;
}
