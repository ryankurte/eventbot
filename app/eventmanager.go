package app

import (
    "time"
)

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

}

// Create an event manager instance
func NewEventManager() *EventManager {

    return &EventManager{}
}

// Create an event
func (em *EventManager) CreateEvent(name string, where string, when time.Time) {
    // Create event instance

    // Send tweet and attach tweet ID to event instance

    // Save to database
}

func (em *EventManager) GetEventByLink(link string) (*Event, error) {

    return nil, nil
}

func (em *EventManager) Update() {
    // Timed update function
}
