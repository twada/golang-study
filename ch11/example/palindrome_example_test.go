package ex

import (
	"github.com/twada/golang-study/ch11/word2"
	"fmt"
)

func ExampleIsPalindrome() {
	fmt.Println(word2.IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(word2.IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}
