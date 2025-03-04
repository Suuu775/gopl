package ex5_7

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func Outline() {
	args := os.Args
	if len(args) > 2 {
		log.Fatal("outline:url")
	}
	resp, err := http.Get(args[1])
	if err != nil {
		return
	}
	n, err := html.Parse(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	forEachNode(n, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s%s/>\n", depth*2, "", n.Data, show_attr(n.Attr))
		} else {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, show_attr(n.Attr))
			depth++
		}
	case html.CommentNode:
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Printf("%*s%s\n", depth*2, "", text)
		}
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild != nil {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

func show_attr(attr []html.Attribute) string {
	var s string
	if len(attr) == 0 {
		s = ""
	} else {
		s = " "
	}
	for _, v := range attr {
		s += fmt.Sprintf("%s=\"%s\" ", v.Key, v.Val)
	}
	return strings.TrimSuffix(s, " ")
}
