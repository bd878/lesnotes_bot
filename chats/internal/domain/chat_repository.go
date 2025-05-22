package domain

import (
	"context"
)

type ChatRepository interface {
	Find(ctx context.Context, chatID int64) (*Chat, error)
	Save(ctx context.Context, chat *Chat) error
}