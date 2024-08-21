package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func normalizeURL(rawURL string) (string, error) {
	rawURL = strings.TrimRight(rawURL, "/")

	url, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	parsedURL := url.Host + url.Path
	return parsedURL, nil
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	rootNode, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}

	// TODO: do stuff with the root node

	fmt.Println("domNodes:", rootNode.FirstChild)
	return []string{}, nil
}
