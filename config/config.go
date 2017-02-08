package config

import (
	"log"
)

import (
	"github.com/jessevdk/go-flags"
	"github.com/kelseyhightower/envconfig"
)

// AuthPlz configuration structure
type EventBotConfig struct {
	Address       string `short:"a" long:"address" description:"Set server address"`
	Port          string `short:"p" long:"port" description:"Set server port"`
	Database      string `short:"d" long:"database" description:"Database connection string"`
	TwitterKey    string `short:"k" long:"twitter-key" description:"Twitter API key"`
	TwitterSecret string `short:"s" long:"twitter-secret" description:"Twitter API secret"`
	TwitterUser   string `short:"u" long:"twitter-user" description:"Twitter username"`
	WatsonUser    string `short:"w" long:"watson-user" description:"Watson API username"`
	WatsonPass    string `short:"x" long:"watson-password" description:"Watson API password"`
	WatsonWs      string `short:"y" long:"watson-workspace" description:"Watson conversations workspace"`
	Test      	  string `short:"t" long:"test" description:"Sets test mode for go test -v ./... support"`
}

// Generate default configuration
func DefaultConfig() (*EventBotConfig, error) {
	var c EventBotConfig

	c.Address = "localhost"
	c.Port = "9000"
	c.Database = "host=localhost user=postgres dbname=postgres sslmode=disable password=postgres"

	return &c, nil
}

func GetConfig() *EventBotConfig {
	// Fetch default configuration
	c, err := DefaultConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Parse config structure through environment
	err = envconfig.Process("EBOT", c)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Override environment with command line args
	_, err = flags.Parse(c)
	if err != nil {
		log.Fatal(err.Error())
	}

	return c
}
