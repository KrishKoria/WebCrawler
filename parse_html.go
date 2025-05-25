package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
    baseURL, err := url.Parse(rawBaseURL)
    if err != nil {
        return nil, err
    }

    htmlReader := strings.NewReader(htmlBody)
    doc, err := html.Parse(htmlReader)
    if err != nil {
        return nil, err
    }

    urls := []string{} 
    extractURLs(doc, baseURL, &urls)
    return urls, nil
}

func extractURLs(n *html.Node, baseURL *url.URL, urls *[]string) {
    if n.Type == html.ElementNode && n.Data == "a" {
        for _, attr := range n.Attr {
            if attr.Key == "href" {
                href := strings.TrimSpace(attr.Val)
                if href != "" {
                    parsedURL, err := url.Parse(href)
                    if err == nil {
                        absoluteURL := baseURL.ResolveReference(parsedURL)
                        *urls = append(*urls, absoluteURL.String())
                    }
                }
                break
            }
        }
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        extractURLs(c, baseURL, urls)
    }
}