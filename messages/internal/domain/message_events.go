package domain

const (
	MessageCreatedEvent = "messages.MessageCreated"
)

type MessageCreated struct {
	Message *Message
}

func (MessageCreated) Key() string { return MessageCreatedEvent }