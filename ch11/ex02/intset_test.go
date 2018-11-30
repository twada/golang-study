package ex02

import (
	"fmt"
	"testing"
)

type SetOfInt interface {
	// Has は負でない値 x をセットが含んでいるか否かを報告します
	Has(x int) bool
	// Add セットに負ではない値 x を追加します
	Add(x int)
}

var table = []struct {
	name string
	body func(SetOfInt) bool
	want bool
}{
	{
		"空の Set",
		func(s SetOfInt) bool {
			return s.Has(3)
		},
		false,
	},
	{
		"Add(7) してから Has(7)",
		func(s SetOfInt) bool {
			s.Add(7)
			return s.Has(7)
		},
		true,
	},
	{
		"Add(-1) してから Has(-1)",
		func(s SetOfInt) bool {
			s.Add(-1)
			return s.Has(-1)
		},
		false,
	},
}

func TestIntSetImpl(t *testing.T) {
	for i, test := range table {
		t.Run(fmt.Sprintf("table#%d : %s", i, test.name), func(t *testing.T) {
			setByMap := NewIntMapSet()
			setByBitVector := &IntSet{}
			got1 := test.body(setByMap)
			got2 := test.body(setByBitVector)
			if got1 != test.want || got2 != test.want {
				t.Errorf("got IntMapSet = %v != IntSet = %v, want %v", got1, got2, test.want)
			}
		})
	}
}

// 以下はベタに書いてみたテスト

func TestHas(t *testing.T) {
	input := 3
	expected := false
	setByMap := NewIntMapSet()
	setByBitVector := &IntSet{}
	actual1 := setByMap.Has(input)
	actual2 := setByBitVector.Has(input)
	if actual1 != expected || actual2 != expected {
		t.Errorf("got IntMapSet.Has(%v) = %v != IntSet.Has(%v) = %v, want %v", input, actual1, input, actual2, expected)
	}
}

func TestAddThenHas(t *testing.T) {
	input := 3
	expected := true
	setByMap := NewIntMapSet()
	setByBitVector := &IntSet{}
	setByMap.Add(input)
	setByBitVector.Add(input)
	actual1 := setByMap.Has(input)
	actual2 := setByBitVector.Has(input)
	if actual1 != expected || actual2 != expected {
		t.Errorf("got IntMapSet.Has(%v) = %v != IntSet.Has(%v) = %v, want %v", input, actual1, input, actual2, expected)
	}
}
