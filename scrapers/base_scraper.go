package scrapers

import (
	"log"

	"github.com/unevenit/scrapie/config"
	"github.com/unevenit/scrapie/utils"
)

type Scraper interface {
	Scrape() (map[string]string, error)
}

type BaseScraper struct {
	config *config.ScrapeConfig
}

func NewScraper(cfg *config.ScrapeConfig) Scraper {
	return &BaseScraper{config: cfg}
}

func (s *BaseScraper) Scrape() (map[string]string, error) {
	log.Printf("Scraping URL: %s", s.config.URL)

	response, err := utils.MakeRequest(s.config.URL, s.config.Method, s.config.Headers)
	if err != nil {
		return nil, err
	}

	log.Printf("Response received: %s", response)

	parsedData, err := utils.ParseResponse(response, s.config.DataSelectors)
	if err != nil {
		return nil, err
	}

	log.Printf("Parsed Data: %v", parsedData)

	return parsedData, nil
}
