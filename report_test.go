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
		"no duplicates per count (number ordering)": {
			pages: map[string]int{
				"https://example.com/page3": 5,
				"https://example.com/page1": 10,
				"https://example.com/page2": 2,
			},
			baseURL: "test.dev.com/path/",
			expected: []reportStruct{
				{Value: 10, Key: "https://example.com/page1"},
				{Value: 5, Key: "https://example.com/page3"},
				{Value: 2, Key: "https://example.com/page2"},
			},
		},
		"one duplicate (alphabetically ordered)": {
			pages: map[string]int{
				"https://example.com/b_key": 5,
				"https://example.com/a_key": 5,
				"https://example.com/c_key": 10,
			},
			baseURL: "test.dev.com/path/",
			expected: []reportStruct{
				{Value: 10, Key: "https://example.com/c_key"},
				{Value: 5, Key: "https://example.com/a_key"}, // 'a_key' comes before 'b_key' alphabetically
				{Value: 5, Key: "https://example.com/b_key"},
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
				{Value: 5, Key: "https://example.com/delta"}, // Same value, delta before gamma
				{Value: 5, Key: "https://example.com/gamma"},
				{Value: 3, Key: "https://example.com/alpha"}, // Same value, alpha before beta
				{Value: 3, Key: "https://example.com/beta"},
				{Value: 2, Key: "https://example.com/epsilon"},
				{Value: 1, Key: "https://example.com/zeta"},
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
