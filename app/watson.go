package app

import (
    "log"
    "time"
)

import(
    "github.com/liviosoares/go-watson-sdk/watson"
    "github.com/liviosoares/go-watson-sdk/watson/conversation"
)

type WatsonConnector struct {
    client *conversation.Client;
    workspace string
}

const(
    watsonConversationUrl string = "https://gateway.watsonplatform.net/conversation/api"
)

func NewWatsonConnector(user string, pass string, workspace string) (*WatsonConnector, error){
    // Create watson configuration
    config := watson.Config{
        Credentials: watson.Credentials{
            Url: watsonConversationUrl,
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
    Response string
    When time.Time
    Location string
}

func (wc* WatsonConnector) HandleMessage(message string) (*WatsonResponse, error) {

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

    resp := WatsonResponse {
        Response: reply.Output.Text[0],
    }

    for _, e := range(reply.Entities) {
        switch e.Entity {
            case "sys-date":

            case "sys-time":

            case "place":
                resp.Location = e.Value
            default:
                log.Printf("Unrecognised conversation entity: %s", e.Entity)
        }
    }

    //intent := reply.Intents[0]

    log.Printf("Response: %+v\n", resp)

    // TODO: return message information
    return &resp, nil
}

