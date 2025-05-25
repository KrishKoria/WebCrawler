package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		println("no website provided")
		os.Exit(1)
	}
	if len(args) > 1 {
		println("too many arguments provided")
		os.Exit(1)
	}
	website := args[0]
	fmt.Printf("starting crawl of: %s\n", website)
	html, err := getHTML(website)
	if err != nil {
		fmt.Printf("error fetching HTML: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("HTML fetched successfully:")
	fmt.Println(html)
}