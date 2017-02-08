/**
 * Generic Message Interface
 * Client messages must implement this interface for binding into the manager
 */

package app

type Message interface {
	// Generate a reply message of the appropriate type
	Reply(text string) interface{}
	// Fetch the text (body) from the message
	Text() string
	// Fetch the user account associated with the message
	User() string
	// Fetch the connector / client associated with the message
	Connector() string
}
