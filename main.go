package main

import (
	// "fmt"
	"flag"
	"github.com/0Blackraven/webscrapper/crawler"
)

func main() {
	var workerCount int;
	var crawlerCount int;
	var url string;
	flag.StringVar(&url, "u", "", "url to be crawled");
	flag.IntVar(&workerCount, "w", 2, "Number of concurrent workers");
	flag.IntVar(&crawlerCount, "c", 1, "Number of concurrent crawlers");

	flag.Parse()
	// fmt.Printf("worker: [%v]\n depth: [%v]\n",workerCount,depthCount);
	crawler.StartCrawling(workerCount, crawlerCount, url);

}