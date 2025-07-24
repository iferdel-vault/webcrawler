package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := map[string]struct {
		html     string
		inputURL string
		expected []string
	}{
		"absolute and relative urls": {
			html: `
			<html>
				<body>
					<a href="/path/one">
						<span>Boot.dev</span>
					</a>
					<a href="https://other.com/path/one">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`,
			inputURL: "https://blog.boot.dev",
			expected: []string{
				"https://blog.boot.dev/path/one",
				"https://other.com/path/one",
			},
		},
		"self contained url": {
			html: `
			<html>
				<body>
					<a href="https://blog.boot.dev/path/one">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`,
			inputURL: "https://blog.boot.dev",
			expected: []string{
				"https://blog.boot.dev/path/one",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := getURLsFromHTML(tc.html, tc.inputURL)
			if err != nil {
				t.Errorf("Test '%s' FAIL: unexpected error: %v", name, err)
				return
			}
			if !reflect.DeepEqual(tc.expected, got) {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}
