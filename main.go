package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
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

	// Hax to keep running
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)
	log.Println("Exiting eventbot")

	// Close server
	server.Close()
}
