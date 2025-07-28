package main

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	tests := map[string]struct {
		pages    map[string]int
		baseURL  string
		expected []reportStruct
	}{
		"no pages": {
			pages:    map[string]int{},
			baseURL:  "test.dev.com/path/",
			expected: []reportStruct{},
		},
		"no duplicates per count (number ordering)": {
			pages: map[string]int{
				"https://example.com/page3": 5,
				"https://example.com/page1": 10,
				"https://example.com/page2": 2,
			},
			baseURL: "test.dev.com/path/",
			expected: []reportStruct{
				{count: 10, url: "https://example.com/page1"},
				{count: 5, url: "https://example.com/page3"},
				{count: 2, url: "https://example.com/page2"},
			},
		},
		"one duplicate (alphabetically ordered)": {
			pages: map[string]int{
				"https://example.com/b_url": 5,
				"https://example.com/a_url": 5,
				"https://example.com/c_url": 10,
			},
			baseURL: "test.dev.com/path/",
			expected: []reportStruct{
				{count: 10, url: "https://example.com/c_url"},
				{count: 5, url: "https://example.com/a_url"}, // 'a_url' comes before 'b_url' alphabetically
				{count: 5, url: "https://example.com/b_url"},
			},
		},
		"random multiple duplicates and not": {
			pages: map[string]int{
				"https://example.com/alpha":   3,
				"https://example.com/beta":    3,
				"https://example.com/zeta":    1,
				"https://example.com/gamma":   5,
				"https://example.com/delta":   5,
				"https://example.com/epsilon": 2,
			},
			baseURL: "test.dev.com/path/",
			expected: []reportStruct{
				{count: 5, url: "https://example.com/delta"}, // Same count, delta before gamma
				{count: 5, url: "https://example.com/gamma"},
				{count: 3, url: "https://example.com/alpha"}, // Same count, alpha before beta
				{count: 3, url: "https://example.com/beta"},
				{count: 2, url: "https://example.com/epsilon"},
				{count: 1, url: "https://example.com/zeta"},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := sortPages(tc.pages, tc.baseURL)
			if err != nil {
				t.Errorf("Test '%s' FAIL: unexpected error: %v", name, err)
				return
			}
			if !reflect.DeepEqual(tc.expected, got) {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
