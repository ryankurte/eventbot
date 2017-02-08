/**
 * Twitter application connector
 *
 */

package clients

import (
	"fmt"
	"log"
)

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"golang.org/x/oauth2"
)

type TwitterConnector struct {
	client *twitter.Client
	stream *twitter.Stream
	ch     chan interface{}
}

// Twitter client message object
type TwitterMessage struct {
	text      string // Message text
	user      string // User name (ie. jpdanner)
	connector string // Connector name (ie. twitter)
	args      string // Arguments for connector (ie. DM)
	userid    int64
	tweetid   int64
}

// Generate a reply message preserving required fields
func (m *TwitterMessage) Reply(text string) interface{} {
	return &TwitterMessage{text, m.user, m.connector, m.args, m.userid, m.tweetid}
}

func (m *TwitterMessage) Text() string {
	return m.text
}

func (m *TwitterMessage) User() string {
	return m.user
}

func (m *TwitterMessage) Connector() string {
	return m.connector
}
func (m *TwitterMessage) Args() string {
	return m.args
}

const (
	TwitterConnectorName string = "twitter"
	twitterPM                   = "public_message"
	twitterDM                   = "direct_message"
)

func NewTwitterConnector(apiKey string, apiSecret string, accessToken string, tokenSecret string, username string) (*TwitterConnector, chan interface{}, error) {

	// Build 2 legged oauth config
	config := oauth1.NewConfig(apiKey, apiSecret)
	token := oauth1.NewToken(accessToken, tokenSecret)

	// OAuth2 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext, token)

	// Twitter client will magically handle authorization
	client := twitter.NewClient(httpClient)

	// Create stream connection
	streamParams := &twitter.StreamUserParams{
		StallWarnings: twitter.Bool(true),
	}
	stream, stream_err := client.Streams.User(streamParams)
	if stream_err != nil {
		log.Printf("Twitter stream error: %s\n", stream_err)
		return nil, nil, stream_err
	}

	ch := make(chan interface{}, 100)

	// Create demux
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		if tweet.User.ScreenName != username {
			log.Printf("Twitter received message: %s from user %s", tweet.Text, tweet.User.ScreenName)
			ch <- &TwitterMessage{tweet.Text, tweet.User.ScreenName, TwitterConnectorName, twitterPM, tweet.User.ID, tweet.ID}
		}
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		if dm.Sender.ScreenName != username {
			log.Printf("Twitter received dm: %s from user %s", dm.Text, dm.Sender.ScreenName)
			ch <- &TwitterMessage{dm.Text, dm.Sender.ScreenName, TwitterConnectorName, twitterDM, dm.Sender.ID, dm.ID}
		}
	}
	demux.Event = func(event *twitter.Event) {
		log.Printf("Twitter received event: %#v", event)
	}

	demux.Other = func(message interface{}) {
		log.Printf("Twitter received unhandled event type: %#v", message)
	}

	// Bind to stream
	go demux.HandleChan(stream.Messages)

	return &TwitterConnector{client, stream, ch}, ch, nil
}

// Send a message using a given connector
func (tc *TwitterConnector) Send(message interface{}) error {
	var err error = nil

	m := message.(*TwitterMessage)

	log.Printf("Twitter: sending message %+v", m)

	switch m.args {
	case twitterDM:
		// Format message
		dm := twitter.DirectMessageNewParams{
			ScreenName: m.User(),
			Text:       fmt.Sprintf("%s", m.Text()),
		}

		// Post DM
		_, _, err = tc.client.DirectMessages.New(&dm)

	case twitterPM:
		// Generate reply metadata
		tweet := twitter.StatusUpdateParams{
			InReplyToStatusID: m.tweetid,
		}

		// Format message
		data := fmt.Sprintf("@%s %s", m.User(), m.Text())

		// Check message length
		if len(data) > 140 {
			return fmt.Errorf("Message length exceeds twitter maximum")
		}

		// Post status update
		_, _, err = tc.client.Statuses.Update(data, &tweet)
	}

	return err
}

func (tc *TwitterConnector) Close() {
	tc.stream.Stop()
}
