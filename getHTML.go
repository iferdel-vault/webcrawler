package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {

	r, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("error on GET request to %q: %s", rawURL, err)
	}
	defer r.Body.Close()

	if r.StatusCode >= 400 {
		return "", fmt.Errorf("status code of respose to GET request on %q is %d", rawURL, r.StatusCode)
	}
	if contentType := r.Header.Get("Content-Type"); !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("response content type to GET request on %q is not 'text/html', is %s", rawURL, contentType)
	}

	html, err := io.ReadAll(r.Body)
	if err != nil {
		return "", fmt.Errorf("error on reading response body: %s", err)
	}

	return string(html), nil
}
