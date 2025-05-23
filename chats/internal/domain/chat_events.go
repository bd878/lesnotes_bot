package domain

const (
	ChatCreatedEvent = "chats.ChatCreated"
	ChatDeletedEvent = "chats.ChatDeleted"
)

type ChatCreated struct {
	Chat *Chat
}

func (ChatCreated) Key() string { return ChatCreatedEvent }

type ChatDeleted struct {
	Chat *Chat
}

func (ChatDeleted) Key() string { return ChatDeletedEvent }