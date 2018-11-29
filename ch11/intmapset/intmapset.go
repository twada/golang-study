package intmapset

type IntMapSet struct {
	ints map[int]struct{}
}

func NewIntMapSet() *IntMapSet {
	s := &IntMapSet{}
	s.ints = make(map[int]struct{})
	return s
}

// Has は負でない値 x をセットが含んでいるか否かを報告します
func (s *IntMapSet) Has(x int) bool {
	if x < 0 {
		return false
	}
	_, ok := s.ints[x]
	return ok
}

// Add セットに負ではない値 x を追加します
func (s *IntMapSet) Add(x int) {
	if x < 0 {
		return
	}
	s.ints[x] = struct{}{}
}
