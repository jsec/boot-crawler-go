package main

import (
	"fmt"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.maxPagesReached() {
		return
	}

	fmt.Println("crawling page:", rawCurrentURL)

	if !hasSameDomain(cfg.baseURL, rawCurrentURL) {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return
	}

	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		return
	}

	urls, err := getURLsFromHTML(htmlBody, normalizedURL)
	if err != nil {
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
