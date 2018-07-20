package main

import (
	"github.com/twada/golang-study/ch06/intset"
	"fmt"
)

func main() {
	x := &intset.IntSet{}
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(1)
	fmt.Println(x)
	fmt.Println(x.Remove(9))
	fmt.Println(x)
}
