package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "USAGE: %s [file1, file2, ...]\n", os.Args[0])
		os.Exit(1)
		return
	}
	for _, arg := range files {
		countLinesInFile(arg, counts)
	}
	for line, dups := range counts {
		n := len(dups)
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, strings.Join(dups, ","))
		}
	}
}

func countLinesInFile(filename string, counts map[string][]string) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2_2: %v\n", err)
		return
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		if (indexOf(counts[text], filename) == -1) {
			counts[text] = append(counts[text], filename)
		}
	}
}

func indexOf(haystack []string, needle string) int {
	for i, entry := range haystack {
		if (entry == needle) {
			return i
		}
	}
	return -1
}
