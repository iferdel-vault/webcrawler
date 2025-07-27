package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]
	fmt.Println("starting crawl of:", rawBaseURL)
	fmt.Println("===============================")

	pages := make(map[string]int)
	crawlPage(rawBaseURL, rawBaseURL, pages)

	for key, value := range pages {
		fmt.Printf("%q: %d\n", key, value)
	}
}
