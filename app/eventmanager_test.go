package app

import (
"log"
"testing"
)

type FakeClient struct {

}

func (fc *FakeClient) Send(m *Message) error {
    log.Printf("Send: %+v\n", m)
    return nil
}

func TestEventManager(t *testing.T) {
    // Fetch configuration
    config := GetConfig()

    // Create watson connector
    wc, err := NewWatsonConnector(config.WatsonUser, config.WatsonPass, config.WatsonWs)
    if err != nil {
        t.Error(err)
        t.FailNow()
    }

    fc := FakeClient{}
    ch := make(chan Message, 100)

    //Create event manager
    em := NewEventManager(nil, wc)

    // Bind test client
    em.BindClient("testclient", &fc, ch)

    // Run tests
    t.Run("Create an event", func(t *testing.T) {
        m := Message{
            "Drinks tonight at Vultures Lane?",
            "testuser",
            "testclient",
            "",
        }

        // Handle message
        err := em.handleMessage(m)
        if err != nil {
            t.Error(err)
            t.FailNow()
        }


    })

    // Tear down user controller

}
