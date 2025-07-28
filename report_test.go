package main

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	tests := map[string]struct {
		pages    map[string]int
		baseURL  string
		expected []ReportStruct
	}{
		"nil map": {
			pages:    nil,
			expected: []ReportStruct{},
		},
		"no duplicates per count (number ordering)": {
			pages: map[string]int{
				"https://example.com/page3": 5,
				"https://example.com/page1": 10,
				"https://example.com/page2": 2,
			},
			baseURL: "test.dev.com/path/",
			expected: []ReportStruct{
				{Count: 10, URL: "https://example.com/page1"},
				{Count: 5, URL: "https://example.com/page3"},
				{Count: 2, URL: "https://example.com/page2"},
			},
		},
		"one duplicate (alphabetically ordered)": {
			pages: map[string]int{
				"https://example.com/b_key": 5,
				"https://example.com/a_key": 5,
				"https://example.com/c_key": 10,
			},
			baseURL: "test.dev.com/path/",
			expected: []ReportStruct{
				{Count: 10, URL: "https://example.com/c_key"},
				{Count: 5, URL: "https://example.com/a_key"}, // 'a_key' comes before 'b_key' alphabetically
				{Count: 5, URL: "https://example.com/b_key"},
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
			expected: []ReportStruct{
				{Count: 5, URL: "https://example.com/delta"}, // Same value, delta before gamma
				{Count: 5, URL: "https://example.com/gamma"},
				{Count: 3, URL: "https://example.com/alpha"}, // Same value, alpha before beta
				{Count: 3, URL: "https://example.com/beta"},
				{Count: 2, URL: "https://example.com/epsilon"},
				{Count: 1, URL: "https://example.com/zeta"},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := sortPages(tc.pages)
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
