package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		return
	}
	rawBaseURL := os.Args[1]

	const maxConcurrency = 3
	const maxPages = 1
	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Println("starting crawl of:", rawBaseURL)
	fmt.Println("===============================")

	cfg.wg.Add(1)
	cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for key, value := range cfg.pages {
		fmt.Printf("%q: %d\n", key, value)
	}
}
