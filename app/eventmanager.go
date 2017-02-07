package app

import (
    "time"
    "log"
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
    Send(m *Message) error
}

// Event instance for storage
type Event struct {
    Name string         // Event name
    Where string        // Event location
    When time.Time      // Event time

    Owner string        // Event owner
    Created time.Time   // Event creation time
    Link string         // Link to event advertisement (ie. Tweet ID)
}

// Event management
type EventManager struct {
    es *EventStore
    clients map[string] ClientConnector
}

// Create an event manager instance
func NewEventManager(es *EventStore) *EventManager {

    clients := make(map[string] ClientConnector)

    return &EventManager{es, clients}
}

func (em *EventManager) BindClient(name string, c ClientConnector, ch chan Message) {
    // Save client instance
    em.clients[name] = c

    // Start thread to listen for client events
    go em.handleMessages(ch)

}

func (em *EventManager) Close() {
    // TODO: exit client routines

}

// Handle messages from a provided channel
func (em *EventManager) handleMessages(c chan Message) {
    for {
        // Load events
        m, open := <- c
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
func (em *EventManager) handleMessage(m Message) {

    // TODO: process incoming message

    // TODO: act on message

    // Generate reply
    reply := m.Reply("Test reply")

    // Locate matching client
    c, ok := em.clients[m.Connector()]
    if !ok {
        log.Printf("Invalid connector %s for response", m.Connector())
    }

    // Send reply
    c.Send(&reply)
}

// Create an event
func (em *EventManager) createEvent(name string, where string, when time.Time) {
    // Create event instance

    // Broadcast and attach ID to event instance

    // Save to database
}


