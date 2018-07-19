package intset

import (
	"testing"
)

func TestHas(t *testing.T) {
	s := &IntSet{}
	actual := s.Has(3)
	expected := false
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}

func TestAddThenHas(t *testing.T) {
	s := &IntSet{}
	s.Add(7)
	actual := s.Has(7)
	expected := true
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}
