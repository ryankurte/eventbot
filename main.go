package main

import (
	"log"
)

import (
	"github.com/jessevdk/go-flags"
	"github.com/kelseyhightower/envconfig"
	"github.com/ryankurte/eventbot/app"
)

func main() {

	log.Println("Starting EventBot")

	// Fetch default configuration
	c, err := app.DefaultConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Parse config structure through environment
	err = envconfig.Process("ebot", c)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Override environment with command line args
	_, err = flags.Parse(c)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create server instance
	server, err := app.NewEventBotServer(c)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Launch server
	server.Start()
}
