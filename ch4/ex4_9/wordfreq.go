package wordfreq

import (
	"bufio"
	"fmt"
	"os"
)

func Wordfreq() {
	wf := make(map[string]uint)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		wf[word]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}

	for k, v := range wf {
		fmt.Printf("freq:%d word:%s", v, k)
	}
}
