// by ai
package main


import (
	"flag"
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/0Blackraven/webscrapper/crawler"
)

func test() {
	var url string
	var workers int
	var crawlers int
	var duration int

	flag.StringVar(&url, "u", "https://books.toscrape.com/", "Target URL for benchmark")
	flag.IntVar(&workers, "w", 10, "Link processing workers")
	flag.IntVar(&crawlers, "c", 10, "HTTP fetching crawlers")
	flag.IntVar(&duration, "d", 30, "Duration in seconds")
	flag.Parse()

	fmt.Printf("\n--- REAL-WORLD BENCHMARK STARTING ---\n")
	fmt.Printf("Target: %s\n", url)
	fmt.Printf("Config: %d fetchers, %d processors\n", crawlers, workers)
	fmt.Printf("Time limit: %ds\n", duration)
	fmt.Println("--------------------------------------")

	crawler.TotalProcessed = 0
	crawler.TotalLinksFound = 0

	startTime := time.Now()

	// Start crawling in a separate goroutine
	go func() {
		crawler.StartCrawling(workers, crawlers, url)
	}()

	// Monitor loop
	timer := time.NewTimer(time.Duration(duration) * time.Second)
	ticker := time.NewTicker(2 * time.Second)
	
loop:
	for {
		select {
		case <-timer.C:
			break loop
		case <-ticker.C:
			proc := atomic.LoadInt64(&crawler.TotalProcessed)
			found := atomic.LoadInt64(&crawler.TotalLinksFound)
			elapsed := time.Since(startTime).Seconds()
			fmt.Printf("Elapsed: %.1fs | Processed: %d (%.2f PPS) | Discovered: %d\n", 
				elapsed, proc, float64(proc)/elapsed, found)
		}
	}

	finalProc := atomic.LoadInt64(&crawler.TotalProcessed)
	finalFound := atomic.LoadInt64(&crawler.TotalLinksFound)
	totalTime := time.Since(startTime).Seconds()

	fmt.Printf("\n--- FINAL RESULTS ---\n")
	fmt.Printf("Total Time: %.2fs\n", totalTime)
	fmt.Printf("Total URLs Processed: %d\n", finalProc)
	fmt.Printf("Total URLs Discovered: %d\n", finalFound)
	fmt.Printf("Throughput: %.2f Pages Per Second\n", float64(finalProc)/totalTime)
	fmt.Println("----------------------")
	
	os.Exit(0)
}
