package main

import (
	"net/url"
	"strings"
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
