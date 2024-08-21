package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[0]

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("invalid maxConcurrency option:", args[1])
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("invalid maxPages option:", args[2])
		os.Exit(1)
	}

	cfg, err := configure(baseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Println("starting crawl of:", baseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, baseURL)
}
