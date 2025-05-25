package main

import (
	"fmt"
	"net/url"
)

func (cfg *Config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done()
	defer func() { <- cfg.concurrencyControl} ()

	cfg.concurrencyControl <- struct{}{}

	cfg.mu.Lock()
    if len(cfg.pages) >= cfg.maxPages {
        cfg.mu.Unlock()
        return
    }
    cfg.mu.Unlock()
    
    currentURL, err := url.Parse(rawCurrentURL)
    if err != nil {
        return
    }
    if cfg.baseURL.Host != currentURL.Host {
        return
    }
    normalizedURL, err := normalizeURL(rawCurrentURL)
    if err != nil {
        return
    }
	isFirst := cfg.addPageVisit(normalizedURL)
    if !isFirst {
        return
    }

    fmt.Printf("crawling: %s\n", rawCurrentURL)
    
    htmlBody, err := getHTML(rawCurrentURL)
    if err != nil {
        fmt.Printf("error getting HTML from %s: %v\n", rawCurrentURL, err)
        return
    }

    urls, err := getURLsFromHTML(htmlBody, cfg.baseURL.String())
    if err != nil {
        fmt.Printf("error parsing URLs from %s: %v\n", rawCurrentURL, err)
        return
    }
    
    for _, foundURL := range urls {
        cfg.wg.Add(1)
		go cfg.crawlPage(foundURL)
    }
}

func (cfg *Config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if count, exists := cfg.pages[normalizedURL]; exists {
		cfg.pages[normalizedURL] = count + 1
		return false
	}
	cfg.pages[normalizedURL] = 1
	return true
}