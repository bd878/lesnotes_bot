package application

import (
	"context"

	"github.com/bd878/lesnotes_bot/internal/ddd"
	"github.com/bd878/lesnotes_bot/messages/internal/domain"
	galleryMessages "github.com/bd878/gallery/server/messages/pkg/model"
)

type (
	CreateMessage struct {
		Message *galleryMessages.Message
		ChatID int64
	}

	App interface {
		CreateMessage(ctx context.Context, cmd CreateMessage) error
	}

	Application struct {
		messages domain.MessagesRepository
		chats domain.ChatRepository
		publisher ddd.EventPublisher[ddd.Event]
	}
)

var _ App = (*Application)(nil)

func New(messages domain.MessagesRepository, chats domain.ChatRepository, publisher ddd.EventPublisher[ddd.Event]) *Application {
	return &Application{
		messages: messages,
		chats: chats,
		publisher: publisher,
	}
}

func (a Application) CreateMessage(ctx context.Context, cmd CreateMessage) error {
	chat, err := a.chats.GetChat(ctx, cmd.ChatID)
	if err != nil {
		return err
	}

	id, err := a.messages.Save(ctx, chat.Token, cmd.Message)
	if err != nil {
		return err
	}

	message, err := domain.CreateMessage(id, cmd.Message, cmd.ChatID)
	if err != nil {
		return err
	}

	return a.publisher.Publish(ctx, message.Events()...)
}
