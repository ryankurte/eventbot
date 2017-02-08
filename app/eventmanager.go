package app

import (
	"fmt"
	"log"
	"time"
)

import (
	// TODO: should probably switch to interfaces to remove these dependencies
	"github.com/ryankurte/eventbot/analysis"
)

// Storage interface for events
type EventStore interface {
}

type User interface {
	GetName() string
	IsAdmin() bool
}

type UserStore interface {
}

// Interface for client connectors
type ClientConnector interface {
	Send(m interface{}) error
}

// Event instance for storage
type Event struct {
	Name  string    // Event name
	Where string    // Event location
	When  time.Time // Event time

	Owner   string    // Event owner
	Created time.Time // Event creation time
	Link    string    // Link to event advertisement (ie. Tweet ID)
}

// Event management
type EventManager struct {
	es      *EventStore
	wc      *analysis.WatsonConnector
	clients map[string]ClientConnector
}

// Create an event manager instance
func NewEventManager(es *EventStore, wc *analysis.WatsonConnector) *EventManager {

	clients := make(map[string]ClientConnector)

	return &EventManager{es, wc, clients}
}

func (em *EventManager) BindClient(name string, c ClientConnector, ch chan interface{}) {
	// Save client instance
	em.clients[name] = c

	// Start thread to listen for client events
	go em.handleMessages(ch)

}

func (em *EventManager) Close() {
	// TODO: exit client routines

}

// Handle messages from a provided channel
func (em *EventManager) handleMessages(c chan interface{}) {
	for {
		// Load events
		m, open := <-c
		if open {
			// Call message handler
			em.handleMessage(m)
		} else {
			// Exit message handling go-routine
			log.Println("Exiting client routine")
			break
		}
	}
}

// Handle a single message
func (em *EventManager) handleMessage(i interface{}) error {

    m, ok := i.(Message)
    if !ok {
        return fmt.Errorf("EventManager message must implement Message interface")
    }

	// TODO: process incoming message
	res, err := em.wc.HandleMessage(m.Text())
	if err != nil {
		return fmt.Errorf("Error processing message")
	}

	// TODO: act on message

	if res.Response != "" {
		// Generate reply
        reply := m.Reply(res.Response)

		// Locate matching client
		c, ok := em.clients[m.Connector()]
		if !ok {
			return fmt.Errorf("Invalid connector %s for response", m.Connector())
		}

		// Send reply
		c.Send(reply)
	}

	return nil
}

// Create an event
func (em *EventManager) createEvent(name string, where string, when time.Time) {
	// Create event instance

	// Broadcast and attach ID to event instance

	// Save to database
}
