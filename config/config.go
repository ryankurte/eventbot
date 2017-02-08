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
	Address            string `short:"a" long:"address" description:"Set server address" split_words:"true"`
	Port               string `short:"p" long:"port" description:"Set server port" split_words:"true"`
	Database           string `short:"d" long:"database" description:"Database connection string" split_words:"true"`
	TwitterApiKey      string `long:"twitter-api-key" description:"Twitter API key" split_words:"true"`
	TwitterApiSecret   string `long:"twitter-api-secret" description:"Twitter API secret" split_words:"true"`
	TwitterAccessToken string `long:"twitter-token" description:"Twitter access token" split_words:"true"`
	TwitterTokenSecret string `long:"twitter-token-secret" description:"Twitter access token secret" split_words:"true"`
	TwitterUser        string `long:"twitter-user" description:"Twitter username" split_words:"true"`
	WatsonUser         string `long:"watson-user" description:"Watson API username" split_words:"true"`
	WatsonPass         string `long:"watson-password" description:"Watson API password" split_words:"true"`
	WatsonWs           string `long:"watson-workspace" description:"Watson conversations workspace" split_words:"true"`
	Test               string `short:"t" long:"test" description:"Sets test mode for go test -v ./... support" split_words:"true"`
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
