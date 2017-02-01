package app

// AuthPlz configuration structure
type EventBotConfig struct {
	Address       string `short:"a" long:"address" description:"Set server address"`
	Port          string `short:"p" long:"port" description:"Set server port"`
	Database      string `short:"d" long:"database" description:"Database connection string"`
	TwitterKey    string `short:"k" long:"twitter-key" description:"Twitter API key"`
	TwitterSecret string `short:"s" long:"twitter-secret" description:"Twitter API secret"`
}

// Generate default configuration
func DefaultConfig() (*EventBotConfig, error) {
	var c EventBotConfig

	c.Address = "localhost"
	c.Port = "9000"
	c.Database = "host=localhost user=postgres dbname=postgres sslmode=disable password=postgres"

	return &c, nil
}
