package crawl2

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Suuu775/gopl/ch5/findlinks2"
)

var tokens = make(chan struct{}, 20)

type linkInfo struct {
	url   string
	depth int
}

func crawl(info linkInfo, maxDepth int) []string {
	if info.depth > maxDepth {
		return nil
	}
	fmt.Println(info.url)
	tokens <- struct{}{}
	list, err := findlinks2.Extract(info.url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func Crawl() {
	maxDepth := 3
	if len(os.Args) > 2 {
		maxDepth, _ = strconv.Atoi(os.Args[2])
	}

	worklist := make(chan []linkInfo)
	unseenLinks := make(chan linkInfo)
	go func() { worklist <- []linkInfo{{url: os.Args[1], depth: 0}} }()

	for i := 0; i < 20; i++ {
		go func() {
			for info := range unseenLinks {
				foundLinks := crawl(info, maxDepth)
				newDepth := info.depth + 1
				for _, link := range foundLinks {
					if newDepth <= maxDepth {
						unseenLinks <- linkInfo{url: link, depth: newDepth}
					}
				}
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, info := range list {
			if !seen[info.url] {
				seen[info.url] = true
				go func(info linkInfo) {
					newList := crawl(info, maxDepth)
					var linkInfos []linkInfo
					newDepth := info.depth + 1
					for _, link := range newList {
						if newDepth <= maxDepth {
							linkInfos = append(linkInfos, linkInfo{url: link, depth: newDepth})
						}
					}
					worklist <- linkInfos
				}(info)
			}
		}
	}
}
