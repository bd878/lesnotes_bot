package chats

import (
	"context"

	"github.com/bd878/lesnotes_bot/internal/system"
	"github.com/bd878/lesnotes_bot/chats/internal/server"
	"github.com/bd878/lesnotes_bot/chats/internal/logging"
	"github.com/bd878/lesnotes_bot/chats/internal/http"
	"github.com/bd878/lesnotes_bot/chats/internal/gateway"
	"github.com/bd878/lesnotes_bot/chats/internal/repository"
	"github.com/bd878/lesnotes_bot/chats/internal/application"
)

type Module struct {}

func (m *Module) Startup(ctx context.Context, mono system.Monolith) error {
	chats := repository.NewChatsRepository("chats.chats", mono.Pool())
	client, err := http.NewClient()
	if err != nil {
		return err
	}
	messages := gateway.NewMessagesGateway(client, mono.Config().MessagesURL)

	app := logging.LogApplicationAccess(application.New(chats, messages), mono.Log())

	if err := server.RegisterBot(app, mono.Bot(), mono.Log()); err != nil {
		return err
	}

	return nil
}

func (m Module) Name() string {
	return "chats"
}