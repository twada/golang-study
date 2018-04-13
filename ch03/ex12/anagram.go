package anagram

import (
	"sort"
)

func isAnagram(s1 string, s2 string) bool {
	return s1 != s2 && sorted(s1) == sorted(s2)
}

func sorted(s string) string {
	runes := make([]rune, len(s)) 
	copy(runes, []rune(s))
	sort.Slice(runes, func(i, j int) bool {
        return runes[i] < runes[j]
    })
	return string(runes)
}
