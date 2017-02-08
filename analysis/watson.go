package analysis

import (
	"log"
)

import (
	"github.com/liviosoares/go-watson-sdk/watson"
	"github.com/liviosoares/go-watson-sdk/watson/conversation"
)

type WatsonConnector struct {
	client    *conversation.Client
	workspace string
}

const (
	watsonConversationUrl string = "https://gateway.watsonplatform.net/conversation/api"
)

func NewWatsonConnector(user string, pass string, workspace string) (*WatsonConnector, error) {
	// Create watson configuration
	config := watson.Config{
		Credentials: watson.Credentials{
			Url:      watsonConversationUrl,
			Username: user,
			Password: pass,
		},
	}

	// Connect to watson service
	client, err := conversation.NewClient(config)
	if err != nil {
		log.Printf("Watson error: %s\n", err)
		return nil, err
	}

	return &WatsonConnector{&client, workspace}, nil
}

type WatsonResponse struct {
	Intent   string
	Response string
	Date     string
	Time     string
	Location string
	Type     string
}

func (wc *WatsonConnector) HandleMessage(message string) (*WatsonResponse, error) {

	// Create watson message
	reply, err := wc.client.Message(wc.workspace, message)
	if err != nil {
		log.Printf("Watson Message() failed %#v\n", err)
		return nil, err
	}

	// Check intent could be inferred
	if len(reply.Intents) == 0 {
		log.Printf("Message() failed.  0 intents returned, expected => 1 intents: %#v\n", reply)
		return nil, nil
	}

	// Create response object
	resp := WatsonResponse{
		Intent: IntentUnrecognized,
	}

	// Update intent if exists
	if len(reply.Intents) > 0 {
		resp.Intent = reply.Intents[0].Intent
	}

	// Save response text if exists
	if len(reply.Output.Text) > 0 {
		resp.Response = reply.Output.Text[0]
	}

	// Parse recognised entities
	for _, e := range reply.Entities {
		switch e.Entity {
		case "sys-date":
			resp.Date = e.Value
		case "sys-time":
			resp.Time = e.Value
		case "place":
			resp.Location = e.Value
		default:
			log.Printf("Unrecognised conversation entity: %s", e.Entity)
		}
	}

	log.Printf("Response: %+v\n", resp)

	// TODO: return message information
	return &resp, nil
}
