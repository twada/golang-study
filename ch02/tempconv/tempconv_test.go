package tempconv

import (
	"testing"
)

func TestConvCToF(t *testing.T) {
	c := Celsius(100)
	actual := CToF(c)
	expected := Fahrenheit(212)
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}
