package main

import (
	"fmt"
	"sort"
)

type ReportStruct struct {
	URL   string
	Count int
}

func printReport(pages map[string]int, baseURL string) {

	fmt.Printf(`
		=============================
			REPORT for %s
		=============================
	`, baseURL)

	sortedResults, err := sortPages(pages)
	if err != nil {
		fmt.Printf("Error sorting pages: %v\n", err)
		return
	}

	for _, page := range sortedResults {
		fmt.Printf("Found %d internal links to %s\n", page.Count, page.URL)
	}

	// flag to save report to csv?
}

func sortPages(pages map[string]int) ([]ReportStruct, error) {

	pagesSlice := []ReportStruct{}
	for url, count := range pages {
		pagesSlice = append(pagesSlice, ReportStruct{
			URL:   url,
			Count: count,
		})
	}

	sort.Slice(pagesSlice, func(i, j int) bool {
		if pagesSlice[i].Count == pagesSlice[j].Count {
			return pagesSlice[i].URL < pagesSlice[j].URL
		}
		return pagesSlice[i].Count > pagesSlice[j].Count
	})

	return pagesSlice, nil
}
