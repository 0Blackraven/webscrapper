package crawler

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/0Blackraven/webscrapper/utils"
)

var TotalProcessed int64
var TotalLinksFound int64

type Worker struct {
	ID int
}

func StartCrawling(workerCount int, crawlerCount int, startURL string) {

	utils.AddJob(startURL, 0)

	linksCh := make(chan string, 100)
	wg := sync.WaitGroup{}

	fmt.Printf("Starting %d workers and %d crawlers...\n", workerCount, crawlerCount)

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for link := range linksCh {
				atomic.AddInt64(&TotalLinksFound, 1)

				utils.AddJob(link, 1)
				fmt.Printf("[Worker %d] Added to queue: %s\n", id, link)
			}
		}(i)
	}

	for i := 1; i <= crawlerCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				url := utils.GetJob()
				if url == "" {
					return
				}

				atomic.AddInt64(&TotalProcessed, 1)
				fmt.Printf("[Crawler %d] Crawling: %s\n", id, url)
				Crawl(url, linksCh)
			}
		}(i)
	}

	wg.Wait()
	close(linksCh)
}
