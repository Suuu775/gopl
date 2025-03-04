package ex52

import "golang.org/x/net/html"

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
