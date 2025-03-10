package ex74

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type StringReader struct {
	s   string
	pos int
}

func NewReader(s string) io.Reader {
	return &StringReader{s: s, pos: 0}
}

func (sr *StringReader) Read(p []byte) (n int, err error) {
	if sr.pos >= len(sr.s) {
		return 0, io.EOF
	}

	n = copy(p, sr.s[sr.pos:])
	sr.pos += n
	return n, nil
}

func Ex74() {
	reader := NewReader("<html><body><h1>Hello, World!</h1></body></html>")

	node, err := html.Parse(reader)
	if err != nil {
		panic(err)
	}

	counts := SameNameElem(node)
	for tag, count := range counts {
		fmt.Printf("%s: %d\n", tag, count)
	}
}

func SameNameElem(n *html.Node) map[string]int {
	var counts = make(map[string]int)
	countElements(n, counts)
	return counts
}

func countElements(n *html.Node, counts map[string]int) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	countElements(n.FirstChild, counts)
	countElements(n.NextSibling, counts)
}
