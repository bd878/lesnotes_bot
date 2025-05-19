package domain

import (
	"context"
)

type MessagesRepository interface {
	Save(ctx context.Context, message *Message) (int32, error)
}