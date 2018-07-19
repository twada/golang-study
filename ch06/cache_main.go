package main

import (
	"github.com/twada/golang-study/ch06/cache"
	"fmt"
)

func main() {
	fmt.Println(cache.Lookup("foo"))
	cache.Set("foo", "hoge")
	fmt.Println(cache.Lookup("foo"))
}
