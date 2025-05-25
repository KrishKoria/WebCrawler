package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type Config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		println("no website provided")
		os.Exit(1)
	}
	if len(args) > 3 {
		println("too many arguments provided")
		os.Exit(1)
	}
	website := args[0]
	maxConcurrencyStr := args[1]
    maxPagesStr := args[2]

	maxConcurrency, err := strconv.Atoi(maxConcurrencyStr)
	if err != nil || maxConcurrency <= 0 {
		fmt.Printf("invalid max concurrency: %s\n", maxConcurrencyStr)
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(maxPagesStr)
	if err != nil || maxPages <= 0 {
		fmt.Printf("invalid max pages: %s\n", maxPagesStr)
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s\n", website)

	baseURL, err := url.Parse(website)
    if err != nil {
        fmt.Printf("error parsing base URL: %v\n", err)
        os.Exit(1)
    }

	cfg := &Config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 10),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(website)
	cfg.wg.Wait()
	
	fmt.Println("\nCrawl complete! Pages found:")
    for url, count := range cfg.pages {
        fmt.Printf("%d: %s\n", count, url)
    }
}