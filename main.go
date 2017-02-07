package main

import (
	"log"
)

import (
	"github.com/ryankurte/eventbot/app"
)

func main() {

	// Fetch configuration
	c := app.GetConfig()

	// Create server instance
	server, err := app.NewEventBotServer(c)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Starting EventBot")

	// Launch server
	server.Start()
}
