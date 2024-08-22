package jobs

import (
	"github.com/unevenit/scrapie/config"
)

type Job struct {
	ID      string
	Config  *config.ScrapeConfig
	Retries int
}

func NewJob(id string, config *config.ScrapeConfig, retries int) *Job {
	return &Job{
		ID:      id,
		Config:  config,
		Retries: retries,
	}
}
