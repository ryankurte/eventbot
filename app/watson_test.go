package app

import (
"testing"
)

func TestWatsonConnector(t *testing.T) {
    // Fetch configuration
    config := GetConfig()

    // Create watson connector
    wc, err := NewWatsonConnector(config.WatsonUser, config.WatsonPass, config.WatsonWs)
    if err != nil {
        t.Error(err)
        t.FailNow()
    }

    // Run tests
    t.Run("Submit a message", func(t *testing.T) {
        _, err := wc.HandleMessage("What's on tonight?")
        if err != nil {
            t.Error(err)
            t.FailNow()
        }
    })

    // Tear down user controller

}
