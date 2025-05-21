package application

import (
	"context"

	"github.com/go-telegram/bot/models"

	"github.com/bd878/lesnotes_bot/chats/internal/domain"
)

type (
	CreateChat struct {
		ID string
		Chat *models.Chat
	}

	CreateMessage struct {
		ID string
		Text string
		UserID int32
	}

	ConfirmIssue struct {
		IssueID int
	}

	KickMember struct {
	}

	App interface {
		CreateChat(ctx context.Context, cmd CreateChat) error
		KickMember(ctx context.Context, cmd KickMember) error
		CreateMessage(ctx context.Context, cmd CreateMessage) (int32, error)
		ConfirmIssue(ctx context.Context, cmd ConfirmIssue) error
	}

	Application struct {
		chats domain.ChatsRepository
		messages domain.MessagesRepository
	}
)

var _ App = (*Application)(nil)

func New(chats domain.ChatsRepository, messages domain.MessagesRepository) *Application {
	return &Application{
		chats: chats,
		messages: messages,
	}
}

func (a Application) CreateChat(ctx context.Context, cmd CreateChat) error {
	chat, err := domain.CreateChat(cmd.ID, cmd.Chat)
	if err != nil {
		return err
	}

	return a.chats.Save(ctx, chat)
}

// TODO: move to separate module "messages"
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

func (a Application) ConfirmIssue(ctx context.Context, cmd ConfirmIssue) error {
	return nil
}

func (a Application) KickMember(ctx context.Context, cmd KickMember) error {
	// TODO: implement
	return nil
}