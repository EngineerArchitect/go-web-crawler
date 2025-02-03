package main

import (
	"fmt"
	"os"
	"strconv"
)

const maxArgs = 4

func main() {
	if len(os.Args) < maxArgs {
		fmt.Println("no website provided")
		return
	}
	if len(os.Args) > maxArgs {
		fmt.Println("too many arguments provided")
		return
	}

	baseURL := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Argument 2 must be int, got %v", err)
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("Argument 3 must be int, got %v", err)
	}

	cfg, err := configure(baseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("error configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", baseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}

	printReport(cfg.pages, baseURL)
}
