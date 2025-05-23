package domain

import (
	"context"
	galleryMessages "github.com/bd878/gallery/server/messages/pkg/model"
)

type MessagesRepository interface {
	Save(ctx context.Context, token string, message *galleryMessages.Message) (int32, error)
}