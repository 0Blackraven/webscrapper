package crawler

import (
	"fmt"
	"strings"

	"github.com/0Blackraven/webscrapper/utils"
)


func Crawl(url string, foundLinks chan<- string) {

	disallowed := utils.RobotResolver(url)

	for _, path := range disallowed {
		if strings.Contains(url, path) {
			fmt.Printf("URL [%s] is disallowed by robots.txt\n", url)
			return
		}
	}

	// utils.HtmlParser(url)
	
	// If HtmlParser is updated to return links, we would send them to the channel:
	links := utils.HtmlParser(url)
	for _, link := range links {
		foundLinks <- link
	}
}
