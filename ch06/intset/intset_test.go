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

func TestString(t *testing.T) {
	s := &IntSet{}
	s.Add(3)
	s.Add(7)
	actual := s.String()
	expected := "{3 7}"
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}

func TestUnionWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	y.Add(9)
	y.Add(42)
	x.UnionWith(&y)
	actual := x.String()
	expected := "{1 9 42 144}"
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}
