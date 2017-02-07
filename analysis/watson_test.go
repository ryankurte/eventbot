package analysis

import (
	"fmt"
	"testing"
)

import (
	"github.com/ryankurte/eventbot/config"
)

func TestWatsonConnector(t *testing.T) {
	// Fetch configuration
	c := config.GetConfig()

	// Create watson connector
	wc, err := NewWatsonConnector(c.WatsonUser, c.WatsonPass, c.WatsonWs)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// Run tests

	t.Run("Create Event intent", func(t *testing.T) {
		resp, err := wc.HandleMessage("Drinks tonight at Vultures Lane?")
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		if resp.Intent != IntentCreateEvent {
			t.Error(fmt.Errorf("Mismatching intents (expected %s received %s)"), IntentCreateEvent, resp.Intent)
		}

	})
	t.Run("Cancel Event", func(t *testing.T) {
		resp, err := wc.HandleMessage("Cancel drinks tonight")
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		if resp.Intent != IntentCancelEvent {
			t.Error(fmt.Errorf("Mismatching intents (expected %s received %s)"), IntentCancelEvent, resp.Intent)
		}

	})
	t.Run("Update event", func(t *testing.T) {
		resp, err := wc.HandleMessage("Move drinks tonight to Brew on Quay")
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		if resp.Intent != IntentUpdateEvent {
			t.Error(fmt.Errorf("Mismatching intents (expected %s received %s)"), IntentUpdateEvent, resp.Intent)
		}

	})
	t.Run("Find Events", func(t *testing.T) {
		resp, err := wc.HandleMessage("What's on tonight?")
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		if resp.Intent != IntentFindEvents {
			t.Error(fmt.Errorf("Mismatching intents (expected %s received %s)"), IntentFindEvents, resp.Intent)
		}

	})
	t.Run("Attend Event", func(t *testing.T) {
		resp, err := wc.HandleMessage("Will be there!")
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		if resp.Intent != IntentAttending {
			t.Error(fmt.Errorf("Mismatching intents (expected %s received %s)"), IntentAttending, resp.Intent)
		}

	})
	t.Run("Not attending event", func(t *testing.T) {
		resp, err := wc.HandleMessage("I can't make it tonight")
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		if resp.Intent != IntentNotAttending {
			t.Error(fmt.Errorf("Mismatching intents (expected %s received %s)"), IntentNotAttending, resp.Intent)
		}

	})
	t.Run("Remind me", func(t *testing.T) {
		resp, err := wc.HandleMessage("Remind me later on")
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		if resp.Intent != IntentRemindMe {
			t.Error(fmt.Errorf("Mismatching intents (expected %s received %s)"), IntentRemindMe, resp.Intent)
		}

	})

	// Tear down user controller

}
