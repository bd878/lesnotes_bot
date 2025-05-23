package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/bd878/lesnotes_bot/chats/chatspb"
	"github.com/bd878/lesnotes_bot/internal/i18n"
	"github.com/bd878/lesnotes_bot/messages/internal/models"
	"github.com/bd878/lesnotes_bot/messages/internal/domain"
)

type ChatRepository struct {
	client chatspb.ChatsServiceClient
}

var _ domain.ChatRepository = (*ChatRepository)(nil)

func NewChatRepository(conn *grpc.ClientConn) ChatRepository {
	return ChatRepository{client: chatspb.NewChatsServiceClient(conn)}
}

func (r ChatRepository) GetChat(ctx context.Context, chatID int64) (*models.Chat, error) {
	resp, err := r.client.GetChat(ctx, &chatspb.GetChatRequest{Id: chatID})
	if err != nil {
		return nil, err
	}

	return r.chatToDomain(resp.Chat), nil
}

func (r ChatRepository) chatToDomain(chat *chatspb.Chat) *models.Chat {
	return &models.Chat{
		ID: chat.Id,
		Token: chat.Token,
		Name: chat.Name,
		Lang: i18n.LangFromString(chat.Lang),
	}
}