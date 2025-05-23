package domain

import (
	"context"
	"github.com/bd878/lesnotes_bot/messages/internal/models"
)

type ChatRepository interface {
	GetChat(ctx context.Context, chatID int64) (*models.Chat, error)
}