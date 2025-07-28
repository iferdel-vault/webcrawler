package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	// Blocks if the channel is full, limiting concurrent goroutines.
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		// Release the a token when the goroutine finishes.
		<-cfg.concurrencyControl
		// Decrement the counter when the goroutine completes.
		cfg.wg.Done()
	}()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("error: couldn't parse current URL:", err)
		return
	}

	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		fmt.Printf("error: baseURL domain is %q and currentURL domain is %q\n", cfg.baseURL.Hostname(), currentURL.Hostname())
		return
	}

	nURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("error: rawCurrentURL cannot be normalized:", err)
		return
	}

	if visited := cfg.addPageVisit(nURL); !visited {
		return
	}

	fmt.Printf("crawling: %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error: getting HTML from %q: %s\n", cfg.baseURL.String(), err)
		return
	}

	urls, err := getURLsFromHTML(html, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("error: getting URL from HTML: %s\n", err)
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go func(url string) {
			cfg.crawlPage(url)
		}(url)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {

	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL] += 1
		return false
	}

	// mark as 'first visited'
	cfg.pages[normalizedURL] = 1
	return true
}
