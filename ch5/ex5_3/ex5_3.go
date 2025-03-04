package ex53

import (
	"golang.org/x/net/html"
)

func TextElemContent(n *html.Node) []string {
	var s []string
	if n == nil {
		return s
	}
	if n.Type == html.TextNode {
		s = append(s, n.Data)
	}
	if n.Data != "style" && n.Data != "script" {
		s = append(s, TextElemContent(n.FirstChild)...)
	}
	s = append(s, TextElemContent(n.NextSibling)...)
	return s
}
