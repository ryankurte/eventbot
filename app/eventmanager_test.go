package app

import (
	"log"
	"testing"
)

import (
	"github.com/ryankurte/eventbot/analysis"
	"github.com/ryankurte/eventbot/config"
)

// Twitter client message object
type FakeMessage struct {
	text      string // Message text
	user      string // User name (ie. jpdanner)
	connector string // Connector name (ie. twitter)
}

// Generate a reply message preserving required fields
func (m *FakeMessage) Reply(text string) interface{} {
	return &FakeMessage{text, m.user, m.connector}
}

func (m *FakeMessage) Text() string {
	return m.text
}

func (m *FakeMessage) User() string {
	return m.user
}

func (m *FakeMessage) Connector() string {
	return m.connector
}

type FakeClient struct {
}

func (fc *FakeClient) Send(m interface{}) error {
    message := m.(*FakeMessage)
	log.Printf("Send: %+v\n", message)
	return nil
}

func (fc *FakeClient) Close() {
    
}

func TestEventManager(t *testing.T) {
	// Fetch configuration
	c := config.GetConfig()

	// Create watson connector
	wc, err := analysis.NewWatsonConnector(c.WatsonUser, c.WatsonPass, c.WatsonWs)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fc := FakeClient{}
	ch := make(chan interface{}, 100)

	//Create event manager
	em := NewEventManager(nil, wc)

	// Bind test client
	em.BindClient("testclient", &fc, ch)

	// Run tests
	t.Run("Handle message directly", func(t *testing.T) {
		m := FakeMessage{
			"Drinks tonight at Vultures Lane?",
			"testuser",
			"testclient",
		}

		// Handle message
		err := em.handleMessage(&m)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	})

    t.Run("Handle message via channel", func(t *testing.T) {
        m := FakeMessage{
            "Drinks tonight at Vultures Lane?",
            "testuser",
            "testclient",
        }

        // Send message via channel
        ch <- m


    })
}
