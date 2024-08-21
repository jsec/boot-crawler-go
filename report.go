package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Println("REPORT for ", baseURL)
	fmt.Println("=============================")

	sortedKeys := getSortedKeys(pages)

	fmt.Println("sortedKeys:", sortedKeys)

	for _, key := range sortedKeys {
		fmt.Println("Found", pages[key], "internal links to", key)
	}
}

func getSortedKeys(pages map[string]int) []string {
	keys := make([]string, 0, len(pages))

	for key := range pages {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return pages[keys[i]] > pages[keys[j]]
	})

	return keys
}
