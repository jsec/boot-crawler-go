package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("Received non-OK response:", err)
	}

	contentType := resp.Header.Get("content-type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("Received invalid content type:", contentType)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response body:", err)
	}
	defer resp.Body.Close()

	return string(body), nil
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	document, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}

	urls := make([]string, 0)

	var inspect func(*html.Node)
	inspect = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" {
					url := a.Val

					firstChar := string([]rune(url)[0])
					if string(firstChar) == "/" {
						url = rawBaseURL + url
					}

					urls = append(urls, url)
				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			inspect(child)
		}
	}

	inspect(document)

	return urls, nil
}
