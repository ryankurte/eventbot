package app

type EventBotServer struct {
    tc *TwitterConnector
    wc *WatsonConnector
}

func NewEventBotServer(config *EventBotConfig) (*EventBotServer, error) {
    var err error

    // Create twitter and Watson API connectors
    tc, err := NewTwitterConnector(config.TwitterKey, config.TwitterSecret, config.TwitterUser)
    if err != nil {
        return nil, err
    }
    wc, err := NewWatsonConnector(config.WatsonUser, config.WatsonPass, config.WatsonWs)
    if err != nil {
        return nil, err
    }

    //TODO: create database connector

    //TODO: create event manager

    //TODO: wire everything together

	return &EventBotServer{tc, wc}, nil
}

func (ebs *EventBotServer) Start() {

}
