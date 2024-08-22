package workers

import (
	"log"

	"github.com/unevenit/scrapie/jobs"
	"github.com/unevenit/scrapie/scrapers"
)

type Worker struct {
	ID string
}

func (w *Worker) Start(jobChannel <-chan *jobs.Job) {
	log.Printf("Worker %s started listening for jobs...", w.ID)

	for job := range jobChannel {
		log.Printf("Worker %s received job %s", w.ID, job.ID)

		scraper := scrapers.NewScraper(job.Config)
		result, err := scraper.Scrape()

		if err != nil {
			log.Printf("Worker %s failed to scrape job %s: %v", w.ID, job.ID, err)
			continue
		}

		log.Printf("Worker %s successfully scraped job %s: %v", w.ID, job.ID, result)
	}

	log.Printf("Worker %s finished processing jobs", w.ID)
}
