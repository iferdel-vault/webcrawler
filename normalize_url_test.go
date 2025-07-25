package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := map[string]struct {
		inputURL string
		expected string
	}{
		"remove scheme": {
			inputURL: "https://test.dev.com/path",
			expected: "test.dev.com/path",
		},
		"remove trailing slash": {
			inputURL: "test.dev.com/path/",
			expected: "test.dev.com/path",
		},
		"remove scheme and trailing slash": {
			inputURL: "http://test.dev.com/path/",
			expected: "test.dev.com/path",
		},
		"remove query parameters": {
			inputURL: "http://test.dev.com/path?test=1245&new=true",
			expected: "test.dev.com/path",
		},
		"remove query parameters again": {
			inputURL: "http://test.dev.com/path?sort=desc",
			expected: "test.dev.com/path",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test '%s' FAIL: unexpected error: %v", name, err)
				return
			}
			if tc.expected != got {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}
