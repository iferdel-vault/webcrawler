package main

import (
	"fmt"
	"log"
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

	html, err := getHTML(rawBaseURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)

}
