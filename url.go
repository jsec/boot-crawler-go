package main

import (
	"log"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	rawURL = strings.TrimRight(rawURL, "/")
	url := parseRawURL(rawURL)

	if url.Scheme == "" {
		url.Scheme = "https"
	}

	return url.String(), nil
}

func hasSameDomain(baseURL, currentURL string) bool {
	base := parseRawURL(baseURL)
	current := parseRawURL(currentURL)

	return base.Hostname() == current.Hostname()
}

func parseRawURL(rawURL string) *url.URL {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		log.Fatal("error parsing URL:", err)
	}

	// Handle URLs without protocol
	if parsedURL.Host == "" {
		parsedURI, err := url.ParseRequestURI("https://" + rawURL)
		if err != nil {
			log.Fatal("error parsing URI:", err)
		}

		return parsedURI
	}

	return parsedURL
}
