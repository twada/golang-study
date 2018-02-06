package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Occurrence struct {
	count int
	files []string
}

func main() {
	occurrences := make(map[string]*Occurrence)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "USAGE: %s [file1, file2, ...]\n", os.Args[0])
		os.Exit(1)
		return
	}
	for _, file := range files {
		countLinesInFile(file, occurrences)
	}
	for line, occ := range occurrences {
		if occ.count > 1 {
			fmt.Printf("%d\t%s\t%v\n", occ.count, line, strings.Join(occ.files, ","))
		}
	}
}

func countLinesInFile(filename string, occurrences map[string]*Occurrence) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2_2: %v\n", err)
		return
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		if _, ok := occurrences[text]; !ok {
			occurrences[text] = new(Occurrence)
		}
		occ := occurrences[text]
		occ.count++
		if (indexOf(filename, occ.files) == -1) {
			occ.files = append(occ.files, filename)
		}
	}
}

func indexOf(needle string, haystack []string) int {
	for i, entry := range haystack {
		if (entry == needle) {
			return i
		}
	}
	return -1
}
