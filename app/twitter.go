/**
 * Twitter application connector
 *
*/

package app

import (
	"fmt"
	"log"
)

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"github.com/dghubble/go-twitter/twitter"
)

type TwitterConnector struct {
	client *twitter.Client
	stream *twitter.Stream
	ch chan Message
}

const (
	TwitterConnectorName string = "twitter"
	twitterPM = "public_message"
	twitterDM = "direct_message"
)

func NewTwitterConnector(apiKey string, apiSecret string, username string) (*TwitterConnector, chan Message, error) {

	// Build 2 legged oauth config
	config := &clientcredentials.Config{
		ClientID:     apiKey,
		ClientSecret: apiSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token"}

	// OAuth2 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)

	// Twitter client will magically handle authorization
	client := twitter.NewClient(httpClient)

	// Fetch a user object
	userShowParams := &twitter.UserShowParams{ScreenName: username}
	user, _, user_err := client.Users.Show(userShowParams)
	if user_err != nil {
		log.Printf("Twitter error: %s\n", user_err)
		return nil, nil, user_err
	}
	fmt.Printf("Got profile for: %s\n", user.ScreenName)

	// Create stream connection
	streamParams := &twitter.StreamUserParams{
		StallWarnings: twitter.Bool(true),
	}
	stream, stream_err := client.Streams.User(streamParams)
	if stream_err != nil {
		log.Printf("Twitter error: %s\n", stream_err)
		return nil, nil, stream_err;
	}

	ch := make(chan Message, 100)

	// Create demux
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		ch <- Message{tweet.Text, tweet.User.Name, TwitterConnectorName, twitterPM}
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		ch <- Message{dm.Text, dm.Sender.Name, TwitterConnectorName, twitterDM}
	}
	demux.Event = func(event *twitter.Event) {
		fmt.Printf("Event: %#v\n", event)
	}

	// Bind to stream
	go demux.HandleChan(stream.Messages)

	return &TwitterConnector{client, stream, ch}, ch, nil
}

// Send a message using a given connector
func (tc *TwitterConnector) Send(message *Message) error {
	// TODO: handle DMs

	// Format message
	data := fmt.Sprintf("@%s %s", message.User(), message.Text());

	// Check message length
	if len(data) > 140 {
		return fmt.Errorf("Message length exceeds twitter maximum")
	}

	// Post status update
	_, _, err := tc.client.Statuses.Update(data, nil)

	return err;
} 

func (tc *TwitterConnector) Close() {
	tc.stream.Stop()
}

