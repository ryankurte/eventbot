package analysis

// User interaction intents
// These map to conversations api intents

type UserIntent string

const (
	IntentUnrecognized string = "Unrecognised" // Unrecognised event
	IntentCreateEvent  string = "CreateEvent"  // Create an event (ie. "Drinks tonight at vultures?")
	IntentCancelEvent  string = "CancelEvent"  // Cancel an event (ie. "Cancel drinks tonight.")
	IntentUpdateEvent  string = "UpdateEvent"  // Update an event (ie. "Move drinks tonight to The Jefferson")
	IntentFindEvents   string = "FindEvents"   // Find events (ie. "What's on tonight")
	IntentAttending    string = "Attending"    // Set attending an event (ie. "I'll be there")
	IntentNotAttending string = "NotAttending" // Set not attending an event (ie. "I'm busy tonight")
	IntentRemindMe     string = "RemindMe"     // Get event reminders (ie. "Remind me the day before")
)
