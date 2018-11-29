package ex

import (
	"fmt"
	"github.com/twada/golang-study/ch11/word2"
)

func ExampleIsPalindrome() {
	fmt.Println(word2.IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(word2.IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}
