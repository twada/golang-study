package popcount

import (
	"testing"
)

func TestPopCountOfSeven(t *testing.T) {
	actual := PopCount(7)
	expected := 3
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}

func TestPopCountOfEight(t *testing.T) {
	actual := PopCount(8)
	expected := 1
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}
