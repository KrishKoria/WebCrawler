package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error){
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		return "", fmt.Errorf("HTTP error: %d %s", res.StatusCode, res.Status)
	}
	contentType := res.Header.Get("Content-Type")
	 if !strings.Contains(strings.ToLower(contentType), "text/html") {
        return "", fmt.Errorf("content-type is not text/html, got: %s", contentType)
    }
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}