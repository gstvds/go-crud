package domain

type Message struct {
	Topic   string      `json:"topic"`
	Version int         `json:"version"`
	Data    interface{} `json:"data"`
}

type MessagingProvider interface {
	Send(topic string, key string, version int, message interface{}, lastCorrelationId ...string)
}
