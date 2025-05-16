package domain

import (
	"context"
)

type ChatsRepository interface {
	FindChat(ctx context.Context, chatID int64) (*Chat, error)
	Save(ctx context.Context, chat *Chat) error
}