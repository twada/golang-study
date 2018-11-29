package randp

import "math/rand"

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // 24 までのランダムな長さ
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // '\u0999' までのランダムなルーン
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}
