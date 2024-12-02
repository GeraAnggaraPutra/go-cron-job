package main

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3" // Importing the cron package for scheduling tasks
)

func main() {
	// Create a new cron scheduler
	c := cron.New()

	// Add a function to be executed every minute
	_, err := c.AddFunc("@every 1m", func() {
		// This function will be called every minute
		fmt.Println("cron job active") // Print a message indicating that the cron job is active
	})
	if err != nil {
		log.Fatal("Error adding function to cron:", err)
	}

	// Start the cron scheduler
	c.Start()

	// Block the main goroutine to keep the program running
	// This prevents the program from exiting immediately
	select {}
}
