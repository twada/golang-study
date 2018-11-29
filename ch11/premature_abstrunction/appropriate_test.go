package premature

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	s, sep := "a:b:c", ":"
	words := strings.Split(s, sep)
	if got, want := len(words), 3; got != want {
		t.Errorf("Split(%q, %q) returned %d words, want %d", s, sep, got, want)
	}
	// ...
}
