package main

import (
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	var urls []string

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}

	urls, err = traverseNodesForURLS(doc)

	for i, url := range urls {
		if !strings.Contains(url, "http") {
			urls[i] = rawBaseURL + url
		} else {
			urls[i] = url
		}
	}

	return urls, nil
}

func traverseNodesForURLS(node *html.Node) ([]string, error) {

	var urls []string

	for n := range node.Descendants() {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "a":
				for _, a := range n.Attr {
					if a.Key == "href" {
						urls = append(urls, a.Val)
					}
				}
			}
		}
	}

	return urls, nil
}

// func traverseNodesForAnchorsRecursive(node *html.Node, urls *[]string) ([]string, error) {
//
// 	for n := range node.Descendants() {
// 		if n.Type == html.ElementNode && n.Data == "a" {
// 			for _, a := range n.Attr {
// 				if a.Key == "href" {
// 					fmt.Printf("found following url in href: %s\n", a.Val)
// 					*urls = append(*urls, a.Val)
// 				}
// 			}
// 		}
// 	}
//
// 	return []string{}, nil
// }
