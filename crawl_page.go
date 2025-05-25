package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
    baseURL, err := url.Parse(rawBaseURL)
    if err != nil {
        return
    }
    
    currentURL, err := url.Parse(rawCurrentURL)
    if err != nil {
        return
    }
    
    if baseURL.Host != currentURL.Host {
        return
    }
    
    normalizedURL, err := normalizeURL(rawCurrentURL)
    if err != nil {
        return
    }
    
    if count, exists := pages[normalizedURL]; exists {
        pages[normalizedURL] = count + 1
        return
    }
    
    pages[normalizedURL] = 1
    
    fmt.Printf("crawling: %s\n", rawCurrentURL)
    
    htmlBody, err := getHTML(rawCurrentURL)
    if err != nil {
        fmt.Printf("error getting HTML from %s: %v\n", rawCurrentURL, err)
        return
    }
    
    urls, err := getURLsFromHTML(htmlBody, rawBaseURL)
    if err != nil {
        fmt.Printf("error parsing URLs from %s: %v\n", rawCurrentURL, err)
        return
    }
    
    for _, foundURL := range urls {
        crawlPage(rawBaseURL, foundURL, pages)
    }
}