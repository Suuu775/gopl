package crawl1

import (
	"fmt"
	"log"
	"os"

	"github.com/Suuu775/gopl/ch5/findlinks2"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := findlinks2.Extract(url)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

func Crawl() {
	worklist := make(chan []string)
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
