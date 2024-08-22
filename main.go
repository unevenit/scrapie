package main

import (
	"log"
	"sync"

	"github.com/unevenit/scrapie/config"
	"github.com/unevenit/scrapie/jobs"
	"github.com/unevenit/scrapie/workers"
)

func main() {
	log.Println("Starting the scraper...")

	configPath := "config.json" // Path to the config file
	scrapeConfig, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Println("Config loaded successfully")

	job := jobs.NewJob("1", scrapeConfig, 3)
	log.Printf("Job created for URL %s", scrapeConfig.URL)

	jobChannel := make(chan *jobs.Job, 1)

	var wg sync.WaitGroup
	wg.Add(1) // Add to the WaitGroup counter

	worker := &workers.Worker{ID: "worker1"}
	log.Println("Starting worker...")
	go func() {
		defer wg.Done() // Signal the WaitGroup when the worker is done
		worker.Start(jobChannel)
	}()

	jobChannel <- job
	log.Println("Job sent to the channel")

	close(jobChannel) // Close the channel once the job is sent

	log.Println("Waiting for worker to finish...")
	wg.Wait() // Wait for the worker to finish

	log.Println("All jobs processed, exiting...")
}
