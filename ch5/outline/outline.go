package outline

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	outline_help(nil, n)
}

func outline_help(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline_help(stack, c)
	}
}
