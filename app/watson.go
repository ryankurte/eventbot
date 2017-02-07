package app

import (
    "log"
)

import(
    "github.com/liviosoares/go-watson-sdk/watson"
    "github.com/liviosoares/go-watson-sdk/watson/conversation"
)

type WatsonConnector struct {
    client *conversation.Client;
    workspace string
}

func NewWatsonConnector(apiKey string, apiSecret string, workspace string) (*WatsonConnector, error){
    // Create watson configuration
    config := watson.Config{
        Credentials: watson.Credentials{
            Username: apiKey,
            Password: apiSecret,
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

func (wc* WatsonConnector) HandleMessage(message string) {

    // Create watson message
    reply, err := wc.client.Message(wc.workspace, message)
    if err != nil {
        log.Printf("Watson Message() failed %#v\n", err)
        return
    }
    
    // Check intent could be inferred
    if len(reply.Intents) == 0 {
        log.Printf("Message() failed.  0 intents returned, expected => 1 intents: %#v\n", reply)
        return
    }

    // TODO: return message information

}

