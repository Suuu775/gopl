package ex55

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (int, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return 0, 0, fmt.Errorf("html Parse:%s", err)
	}
	words, images := countWordsAndImages(doc)
	return words, images, nil
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	if n.Type == html.TextNode {
		sArr := strings.Split(n.Data, " ")
		words += len(sArr)
	} else if n.Data == "img" {
		images = 1
	}
	var wc, ic int
	if n.Data != "style" && n.Data != "script" {
		wc, ic = countWordsAndImages(n.FirstChild)
	}
	wb, ib := countWordsAndImages(n.NextSibling)
	words = words + wc + wb
	images = images + ib + ic
	return
}
