package premature

import (
	"fmt"
	"strings"
	"testing"
)

// 出来が悪いアサート関数
func assertEqual(x, y int) {
	if x != y {
		panic(fmt.Sprintf("%d != %d", x, y))
	}
}

func TestSplit(t *testing.T) {
	words := strings.Split("a:b:c", ":")
	assertEqual(len(words), 3)
	// ...
}
