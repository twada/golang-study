package intset

import (
	"fmt"
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

func TestLen(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	y.Add(9)
	y.Add(42)
	actual := x.Len()
	expected := 3
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
	actual = y.Len()
	expected = 2
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}

func TestRemove(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	actual := x.Remove(144)
	expected := true
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
	actualStr := x.String()
	expectedStr := "{1 9}"
    if actualStr != expectedStr {
        t.Errorf("got %v\nwant %v", actualStr, expectedStr)
    }
}

func TestClear(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	orig := x.words
	actual := fmt.Sprintf("%v", orig)
	expected := "[514 0 65536]"
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
	x.Clear()
	actual = x.String()
	expected = "{}"
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
	actual = fmt.Sprintf("%v", x.words)
	expected = "[]"
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
	actual = fmt.Sprintf("%v", orig)
	expected = "[514 0 65536]"
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}

func TestCopy(t *testing.T) {
	var orig IntSet
	orig.Add(1)
	orig.Add(144)
	orig.Add(9)
	actual := orig.String()
	expected := "{1 9 144}"
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
	cp := orig.Copy()
	actual = cp.String()
	expected = "{1 9 144}"
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
	cp.Add(4)
	cp.Add(6)
	actual = cp.String()
	expected = "{1 4 6 9 144}"
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
	actual = orig.String()
	expected = "{1 9 144}"
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}

func TestAddAll(t *testing.T) {
	s := &IntSet{}
	s.AddAll(3, 6, 9, 12)
	actual := s.String()
	expected := "{3 6 9 12}"
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}

func TestElems(t *testing.T) {
	s := &IntSet{}
	s.AddAll(3, 6, 9, 12)
	var actual string
	for _, x := range s.Elems() {
		actual += fmt.Sprintf(" %d", x)
	}
	expected := " 3 6 9 12"
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}
