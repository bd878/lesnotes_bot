package domain

import (
	"context"
)

type ChatRepository interface {
	Load(ctx context.Context, chatID int64) (*Chat, error)
	UpdateToken(ctx context.Context, chatID int64, token string) error
	Save(ctx context.Context, chat *Chat) error
	Remove(ctx context.Context, chatID int64) error
}