package main

import "fmt"

type reportStruct struct {
	count int
	url   string
}

func printReport(pages map[string]int, baseURL string) {

	fmt.Println(`
		=============================
			REPORT for https://example.com
		=============================
	`)

	sortedResults, err := sortPages(pages, baseURL)
	if err != nil {
		fmt.Printf("Error sorting pages: %v\n", err)
		return
	}

	for _, r := range sortedResults {
		fmt.Printf("Count: %d, URL: %s\n", r.count, r.url)
	}

	// flag to save report to csv?
}

func sortPages(pages map[string]int, baseURL string) ([]reportStruct, error) {
	var results []reportStruct

	return results, nil
}
