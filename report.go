package main

import (
	"fmt"
	"sort"
)

type Page struct {
    URL   string
    Count int
}

func printReport(pages map[string]int, baseURL string) {
    fmt.Printf("=============================\n")
    fmt.Printf("  REPORT for %s\n", baseURL)
    fmt.Printf("=============================\n")

    sortedPages := sortPages(pages)
    
    for _, page := range sortedPages {
        fmt.Printf("Found %d internal links to %s\n", page.Count, page.URL)
    }
}

func sortPages(pages map[string]int) []Page {
    pageSlice := make([]Page, 0, len(pages))
    for url, count := range pages {
        pageSlice = append(pageSlice, Page{URL: url, Count: count})
    }

    sort.Slice(pageSlice, func(i, j int) bool {
        if pageSlice[i].Count == pageSlice[j].Count {
            return pageSlice[i].URL < pageSlice[j].URL
        }
        return pageSlice[i].Count > pageSlice[j].Count
    })

    return pageSlice
}