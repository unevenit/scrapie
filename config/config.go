package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/unevenit/scrapie/utils"
)

type ScrapeConfig struct {
	URL           string                        `json:"url"`
	Method        string                        `json:"method"`
	Headers       map[string]string             `json:"headers"`
	DataSelectors map[string]utils.DataSelector `json:"data_selectors"`
}

func LoadConfig(filePath string) (*ScrapeConfig, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config ScrapeConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
