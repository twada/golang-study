package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("index: %d value: %s\n", i, arg)
	}
}
