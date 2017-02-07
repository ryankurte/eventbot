/**
 * Generic Message
 * Used for communication between services
*/

package app

type Message struct {
    text string         // Message text
    user string        // User name (ie. jpdanner)
    connector string    // Connector name (ie. twitter)
    args string         // Arguments for connector (ie. DM)
};

// Generate a reply message preserving required fields
func (m *Message) Reply(text string) Message {
    return Message{text, m.user, m.connector, m.args}
}

// Methods to fulfil interface requirements for later modules

func (m *Message) Text() string {
    return m.text
}

func (m *Message) User() string {
    return m.user
}

func (m *Message) Connector() string {
    return m.connector
}
func (m *Message) Args() string {
    return m.args
}


