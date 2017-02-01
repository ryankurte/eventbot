package app

// User interaction intents
// These will approximately map to conversations api intents
const (
    IntentCreateEvent = iota    // Create an event (ie. "Drinks tonight at vultures?")
    IntentCancelEvent = iota    // Cancel an event (ie. "Cancel drinks tonight.")
    IntentUpdateEvent = iota    // Update an event (ie. "Move drinks tonight to The Jefferson")
    IntentFindEvents = iota     // Find events (ie. "What's on tonight")
    IntentSubscribe = iota      // Subscribe to events (ie. "Let me know what's happening")
    IntentUnsubscribe = iota    // Subscribe to events (ie. "Unsubscribe")
    IntentRemindMe = iota       // Get event reminders (ie. "Remind me the day before")
)
