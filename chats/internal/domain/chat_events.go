package domain

const (
	ChatCreatedEvent = "chats.ChatCreated"
)

type ChatCreated struct {
	Chat *Chat
}

func (ChatCreated) Key() string { return ChatCreatedEvent }