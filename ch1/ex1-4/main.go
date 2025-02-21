package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int) // map[line]map[filename]count
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}
	for line, fileCounts := range counts {
		for filename, count := range fileCounts {
			if count > 1 {
				fmt.Printf("%d\t%s\t%s\n", count, filename, line)
			}
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if _, ok := counts[line]; !ok {
			counts[line] = make(map[string]int)
		}
		counts[line][filename]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
