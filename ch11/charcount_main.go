// charcount は Unicode 文字の数を計算します
package main

import (
	"fmt"
	"github.com/twada/golang-study/ch11/charcount"
	"os"
)

func main() {
	counts, utflen, err := charcount.Count(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("rune\tconut\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tconut\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
}
