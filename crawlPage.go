package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("error: couldn't parse base URL:", err)
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("error: couldn't parse current URL:", err)
		return
	}

	if baseURL.Hostname() != currentURL.Hostname() {
		fmt.Printf("error: baseURL domain is %q and currentURL domain is %q\n", baseURL.Hostname(), currentURL.Hostname())
		return
	}

	nURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("error: rawCurrentURL cannot be normalized:", err)
		return
	}

	if _, visited := pages[nURL]; visited {
		pages[nURL] += 1
		return
	}

	// mark as 'first visited'
	pages[nURL] = 1

	fmt.Printf("crawling: %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error: getting HTML from %q: %s\n", rawBaseURL, err)
		return
	}

	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		fmt.Printf("error: getting URL from HTML: %s\n", err)
		return
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
