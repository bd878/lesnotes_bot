package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/bd878/lesnotes_bot/chats/chatspb"
	"github.com/bd878/lesnotes_bot/chats/internal/application"
	"github.com/bd878/lesnotes_bot/chats/internal/domain"
)

type server struct {
	app application.App
	chatspb.UnimplementedChatsServiceServer
}

var _ chatspb.ChatsServiceServer = (*server)(nil)

func RegisterServer(_ context.Context, app application.App, registrar *grpc.Server) error {
	chatspb.RegisterChatsServiceServer(registrar, server{app: app})
	return nil
}

func (s server) GetChat(ctx context.Context, req *chatspb.GetChatRequest) (*chatspb.GetChatResponse, error) {
	chat, err := s.app.GetChat(ctx, application.GetChat{ID: req.Id})
	if err != nil {
		return nil, err
	}

	return &chatspb.GetChatResponse{
		Chat: domainToProto(chat),
	}, nil
}

func domainToProto(chat *domain.Chat) *chatspb.Chat {
	return &chatspb.Chat{
		Id: chat.Chat.ID,
		Token: chat.Token,
		Name: chat.Login,
		Lang: chat.Lang.String(),
	}
}