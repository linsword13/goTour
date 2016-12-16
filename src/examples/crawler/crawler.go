// Before using, please make sure you have
// golang.org/x/net/html package installed
// If not, run 'go get golang.org/x/net/html'
package main

import (
	"./links"
	"flag"
	"fmt"
	"log"
)

var (
	startLink = flag.String("start", "https://google.com", "starting link")
	maxWorker = flag.Int("n", 10, "maximum amount of concurrent workers")
	maxDepth  = flag.Int("d", 1, "maximum depth of crawling")
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	flag.Parse()
	mWorker := *maxWorker
	mDepth := *maxDepth
	startLinks := []string{*startLink}
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() { worklist <- startLinks }()

	for i := 0; i < mWorker; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	linkNOld := 0
	linkNNew := len(startLinks)
	for depth := 0; depth <= mDepth+1; depth++ {
		linkNOld, linkNNew = linkNNew, 0
		linkList := make([]string, 0)
		for list := range worklist {
			linkNOld--
			if depth == mDepth+1 {
				if linkNOld <= 0 {
					break
				}
				continue
			}
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					linkList = append(linkList, link)
					linkNNew++
				}
			}
			if linkNOld <= 0 {
				for _, newLink := range linkList {
					unseenLinks <- newLink
				}
				break
			}
		}
	}
}
