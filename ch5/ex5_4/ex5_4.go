package ex54

import "golang.org/x/net/html"

func Visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		case "img":
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					links = append(links, attr.Val)
				}
			}
		case "script":
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					links = append(links, attr.Val)
				}
			}
		case "link":
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
	}

	// Traverse child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = Visit(links, c)
	}
	return links
}
