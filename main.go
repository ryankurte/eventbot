package main

import (
	"log"
)

import (
	"github.com/ryankurte/eventbot/app"
	"github.com/ryankurte/eventbot/config"
)

func main() {

	// Fetch configuration
	c := config.GetConfig()

	// Create server instance
	server, err := app.NewEventBotServer(c)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Starting EventBot")

	// Launch server
	server.Start()
}
