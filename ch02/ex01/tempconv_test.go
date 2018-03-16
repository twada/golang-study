package tempconv

import (
	"testing"
)

func TestConvCToK(t *testing.T) {
	actual := CToK(BoilingC).String()
	expected := "373.15K"
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}
