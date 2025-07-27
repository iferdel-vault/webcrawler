package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("couln't parse html: %v", err)
	}

	var urls []string
	err = traverseNodesForURLSRecursively(doc, *baseURL, &urls)
	// urls, err = traverseNodesForURLS(doc, *baseURL)
	return urls, nil
}

func traverseNodesForURLSRecursively(node *html.Node, baseURL url.URL, urls *[]string) error {

	if node.Type == html.ElementNode && node.Data == "a" {
		for _, anchor := range node.Attr {
			href, err := url.Parse(anchor.Val)
			if err != nil {
				fmt.Printf("couldn't parse href '%v': %v\n", anchor.Val, err)
				continue
			}
			resolvedURL := baseURL.ResolveReference(href)
			*urls = append(*urls, resolvedURL.String())
		}
	}
	// for each children(element) check if anchor and then run the recursion
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		traverseNodesForURLSRecursively(child, baseURL, urls)
	}

	return nil
}

func traverseNodesForURLS(node *html.Node, baseURL url.URL) ([]string, error) {

	var urls []string

	for n := range node.Descendants() {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "a":
				for _, a := range n.Attr {
					if a.Key == "href" {
						href, err := url.Parse(a.Val)
						if err != nil {
							fmt.Printf("couldn't parse href '%v': %v\n", a.Val, err)
							continue
						}

						resolvedURL := baseURL.ResolveReference(href)
						urls = append(urls, resolvedURL.String())
					}
				}
			}
		}
	}

	return urls, nil
}
