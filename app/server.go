package app

type EventBotServer struct {
    tc *TwitterConnector
    wc *WatsonConnector
    em *EventManager
}

func NewEventBotServer(config *EventBotConfig) (*EventBotServer, error) {
    var err error

    //TODO: create database connector

    // Create watson connector
    wc, err := NewWatsonConnector(config.WatsonUser, config.WatsonPass, config.WatsonWs)
    if err != nil {
        return nil, err
    }

    //Create event manager
    em := NewEventManager(nil)

    // Create twitter API client
    tc, ch, err := NewTwitterConnector(config.TwitterKey, config.TwitterSecret, config.TwitterUser)
    if err != nil {
        return nil, err
    }

    // Bind to Event manager
    em.BindClient(TwitterConnectorName, tc, ch)

	return &EventBotServer{tc, wc, em}, nil
}

func (ebs *EventBotServer) Start() {
    // TODO: start, maybe?
    // Could have a shitty web view here also

}
