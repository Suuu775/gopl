package ex71

import (
	"bufio"
	"bytes"
)

type WordLineCounter struct {
	wordCount int
	lineCount int
}

func (c *WordLineCounter) Write(p []byte) (int, error) {

	buf := bytes.NewBuffer(p)
	scanner := bufio.NewScanner(buf)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		c.lineCount++
		line := scanner.Text()

		wordScanner := bufio.NewScanner(bytes.NewBuffer([]byte(line)))
		wordScanner.Split(bufio.ScanWords)
		wordCount := 0
		for wordScanner.Scan() {
			wordCount++
		}
		c.wordCount += wordCount
	}

	return len(p), nil
}
