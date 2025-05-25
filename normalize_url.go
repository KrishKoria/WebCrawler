package main

import (
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
    parsedURL, err := url.Parse(inputURL)
    if err != nil {
        return "", err
    }

    host := strings.ToLower(parsedURL.Host)
    
    path := parsedURL.Path
    if len(path) > 1 && strings.HasSuffix(path, "/") {
        path = strings.TrimSuffix(path, "/")
    }
    
    normalized := host + path
    
    if parsedURL.RawQuery != "" {
        normalized += "?" + parsedURL.RawQuery
    }
    
    if parsedURL.Fragment != "" {
        normalized += "#" + parsedURL.Fragment
    }
    
    return normalized, nil
}