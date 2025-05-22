package application

import (
	"context"

	"github.com/bd878/lesnotes_bot/messages/internal/domain"
)

type (
	CreateMessage struct {
		ID string
		Text string
		UserID int32
	}

	App interface {
		CreateMessage(ctx context.Context, cmd CreateMessage) (int32, error)
	}

	Application struct {
		messages domain.MessagesRepository
	}
)

var _ App = (*Application)(nil)

func New(messages domain.MessagesRepository) *Application {
	return &Application{
		messages: messages,
	}
}

func (a Application) CreateMessage(ctx context.Context, cmd CreateMessage) (int32, error) {
	msg, err := domain.CreateMessage(cmd.ID, cmd.Text, cmd.UserID)
	if err != nil {
		return 0, err
	}
	id, err := a.messages.Save(ctx, msg)
	if err != nil {
		return 0, err
	}
	return id, err
}
