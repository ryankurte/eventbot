package app

//import "os"

//import "strings"
import "fmt"
import "log"
import "flag"
import "os"
import "time"

import "golang.org/x/oauth2"
import "golang.org/x/oauth2/clientcredentials"
import "github.com/dghubble/go-twitter/twitter"
import "github.com/coreos/pkg/flagutil"

type TwitterConnector struct {
	client *twitter.Client
}

func (tc *TwitterConnector) NewTwitterConnector(apiKey string, apiSecret string) {

	// Build 2 legged oauth config
	config := &clientcredentials.Config{
		ClientID:     *apiKey,
		ClientSecret: *apiSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token"}

	// OAuth2 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)

	// Twitter client will magically handle authorization
	client := twitter.NewClient(httpClient)

	// Fetch a user object
	userShowParams := &twitter.UserShowParams{ScreenName: "ryankurte"}
	user, _, user_err := client.Users.Show(userShowParams)
	if user_err != nil {
		log.Fatal(user_err)
	}
	fmt.Printf("Got profile for: %s\n", user.ScreenName)

	return TwitterConnector{client}
}
