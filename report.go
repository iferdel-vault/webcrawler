package main

import (
	"fmt"
	"sort"
)

type reportStruct struct {
	Key   string
	Value int
}

func printReport(pages map[string]int, baseURL string) {

	fmt.Printf(`
		=============================
			REPORT for %s
		=============================
	`, baseURL)

	sortedResults, err := sortPages(pages, baseURL)
	if err != nil {
		fmt.Printf("Error sorting pages: %v\n", err)
		return
	}

	for _, r := range sortedResults {
		fmt.Printf("Found %d internal links to %s\n", r.Value, r.Key)
	}

	// flag to save report to csv?
}

func sortPages(pages map[string]int, baseURL string) ([]reportStruct, error) {

	var ss []reportStruct
	for key, value := range pages {
		ss = append(ss, reportStruct{key, value})
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value == ss[j].Value {
			return ss[i].Key < ss[j].Key
		}
		return ss[i].Value > ss[j].Value
	})

	return ss, nil
}
