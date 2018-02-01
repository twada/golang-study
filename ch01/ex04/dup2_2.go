package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	counts := make(map[string]int)
	if len(args) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for filename := range args {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2_2: %v\n", err)
				continue
			}
			countLines(f, counts)
		}
	}
}

func conutLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
