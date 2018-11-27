// charcount は Unicode 文字の数を計算します
package charcount

import (
	"bufio"
	"io"
	"unicode"
	"unicode/utf8"
)

func Count(input io.Reader) (counts map[rune]int, utflen []int, err error) {
	counts = make(map[rune]int)
	utflen = make([]int, utf8.UTFMax + 1)
	invalid := 0

	in := bufio.NewReader(input)
	for {
		r, n, err := in.ReadRune() // rune, nbytes, error を返す
		if err == io.EOF {
			break
		}
		if err != nil {
			return counts, utflen, err
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	return counts, utflen, nil
}
