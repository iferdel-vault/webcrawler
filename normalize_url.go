package main

import (
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	url, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}
	urlPathNoLeadingSlash, _ := strings.CutSuffix(url.Path, "/")
	normalizedURL := url.Host + urlPathNoLeadingSlash
	return normalizedURL, nil
}
