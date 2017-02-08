package app

import (
	"log"
	"sync"
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
	Close()
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
	es       *EventStore
	wc       *analysis.WatsonConnector
	clients  map[string]ClientConnector
	channels map[string]chan interface{}
	wg       sync.WaitGroup
}

// Create an event manager instance
func NewEventManager(es *EventStore, wc *analysis.WatsonConnector) *EventManager {

	clients := make(map[string]ClientConnector)
	channels := make(map[string]chan interface{})

	return &EventManager{es, wc, clients, channels, sync.WaitGroup{}}
}

func (em *EventManager) BindClient(name string, c ClientConnector, ch chan interface{}) {
	// Save client instance
	em.clients[name] = c
	em.channels[name] = ch

	// Start thread to listen for client events
	em.wg.Add(1)
	go em.handleMessages(ch)

}

func (em *EventManager) Close() {
	// TODO: exit client routines
	for _, c := range em.channels {
		close(c)
	}

	em.wg.Wait()
}

// Handle messages from a provided channel
func (em *EventManager) handleMessages(c chan interface{}) {
	log.Printf("Starting message handler routine")

	for {
		// Load events
		m, open := <-c
		if open {
			// Call message handler
			log.Printf("Received message generic")
			em.handleMessage(m)
		} else {
			// Exit message handling go-routine
			log.Println("Exiting client routine")
			break
		}
	}
	log.Printf("Exiting message handler routine")
	em.wg.Done()
}

// Handle a single message
func (em *EventManager) handleMessage(i interface{}) error {
	var err error = nil

	m := i.(Message)

	log.Printf("Processing message")

	// TODO: process incoming message
	res, err := em.wc.HandleMessage(m.Text())
	if err != nil {
		log.Printf("Error processing message")
		return err
	}

	// TODO: act on message
	switch res.Intent {
	case analysis.IntentCreateEvent:

	case analysis.IntentCancelEvent:

	case analysis.IntentUpdateEvent:

	case analysis.IntentFindEvents:

	case analysis.IntentAttending:

	case analysis.IntentNotAttending:

	case analysis.IntentRemindMe:

	}

	if res.Response != "" {
		log.Printf("Sending reply")

		// Generate reply
		reply := m.Reply(res.Response)

		// Locate matching client
		c, ok := em.clients[m.Connector()]
		if !ok {
			log.Printf("Invalid connector %s for response", m.Connector())
			return err
		}

		// Send reply
		c.Send(reply)
	}

	return err
}

// Create an event
func (em *EventManager) createEvent(name string, where string, when time.Time) {
	// Create event instance

	// Broadcast and attach ID to event instance

	// Save to database
}
