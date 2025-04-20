package chats

import (
	"context"

	"github.com/bd878/lesnotes_bot/internal/system"
	"github.com/bd878/lesnotes_bot/chats/internal/server"
	"github.com/bd878/lesnotes_bot/chats/internal/logging"
	"github.com/bd878/lesnotes_bot/chats/internal/repository"
	"github.com/bd878/lesnotes_bot/chats/internal/application"
)

type Module struct {}

func (m *Module) Startup(ctx context.Context, mono system.Monolith) error {
	chats := repository.NewChatsRepository("chats.chats", mono.Pool())

	app := logging.LogApplicationAccess(application.New(chats, mono.Log()), mono.Log())

	if err := server.RegisterBot(app, mono.Bot()); err != nil {
		return err
	}

	return nil
}

func (m Module) Name() string {
	return "chats"
}