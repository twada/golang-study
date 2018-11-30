package ex07

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

type SetOfInt interface {
	// Has は負でない値 x をセットが含んでいるか否かを報告します
	Has(x int) bool
	// Add セットに負ではない値 x を追加します
	Add(x int)
}

var table = []struct {
	name   string
	create func() SetOfInt
}{
	{"BitVector", func() SetOfInt { return &IntSet{} }},
	{"map", func() SetOfInt { return NewIntMapSet() }},
}

func TestIntSet(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	inputs := make([]int, 5)
	for i := 0; i < 5; i++ {
		inputs[i] = rng.Intn(math.MaxInt64) / 1000000000
	}

	for _, impl := range table {
		for _, input := range inputs {
			t.Run(fmt.Sprintf("%v/Add(%d)", impl.name, input), func(t *testing.T) {
				set := impl.create()
				set.Add(input)
				if got := set.Has(input); got != true {
					t.Errorf("%v.Has(%d) = false", impl.name, input)
				}
			})
		}
	}
}

func BenchmarkIntSet(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	inputs := make([]int, 2)
	for i := 0; i < 2; i++ {
		inputs[i] = rng.Intn(math.MaxInt64) / 100000000
		// inputs[i] = rng.Intn(math.MaxInt32)
	}

	for _, input := range inputs {
		for _, impl := range table {
			b.Run(fmt.Sprintf("%v/Add(%d)", impl.name, input), func(b *testing.B) {
				set := impl.create()
				for i := 0; i < b.N; i++ {
					set.Add(input)
				}
			})
		}
	}
}
