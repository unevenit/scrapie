package utils

import (
	"errors"
	"strings"

	"golang.org/x/net/html"
)

// ParseResponse now extracts the actual title from the HTML
func ParseResponse(response string, selectors map[string]string) (map[string]string, error) {
	parsedData := make(map[string]string)

	doc, err := html.Parse(strings.NewReader(response))
	if err != nil {
		return nil, err
	}

	for key, selector := range selectors {
		if selector == "title" {
			parsedData[key] = extractTitle(doc)
		} else {
			return nil, errors.New("Unsupported selector")
		}
	}

	return parsedData, nil
}

// extractTitle navigates the HTML tree to find the <title> tag
func extractTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" {
		if n.FirstChild != nil {
			return n.FirstChild.Data
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := extractTitle(c)
		if result != "" {
			return result
		}
	}
	return ""
}
