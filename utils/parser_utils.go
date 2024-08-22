package utils

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// DataSelector represents the structure of each data extraction option in the config.
type DataSelector struct {
	Selector  string `json:"selector"`
	Type      string `json:"type"`
	Attribute string `json:"attribute,omitempty"` // Optional, used only for "attribute" type
}

// ParseResponse parses the HTML response and extracts data based on the provided selectors.
func ParseResponse(response string, selectors map[string]DataSelector) (map[string]string, error) {
	parsedData := make(map[string]string)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		return nil, err
	}

	// Loop through the selectors and extract data based on the type
	for key, selectorConfig := range selectors {
		switch selectorConfig.Type {
		case "text":
			parsedData[key] = doc.Find(selectorConfig.Selector).First().Text()
		case "attribute":
			parsedData[key], _ = doc.Find(selectorConfig.Selector).First().Attr(selectorConfig.Attribute)
		case "texts":
			var texts []string
			doc.Find(selectorConfig.Selector).Each(func(i int, s *goquery.Selection) {
				texts = append(texts, s.Text())
			})
			parsedData[key] = strings.Join(texts, "\n")
		default:
			return nil, fmt.Errorf("unsupported selector type: %s", selectorConfig.Type)
		}
	}

	return parsedData, nil
}
