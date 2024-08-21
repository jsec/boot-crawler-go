package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	rootNode, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}

	// TODO: do stuff with the root node

	fmt.Println("domNodes:", rootNode.FirstChild)
	return []string{}, nil
}
