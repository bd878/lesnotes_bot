package chats

import (
	"context"

	"github.com/bd878/lesnotes_bot/internal/system"
	"github.com/bd878/lesnotes_bot/internal/ddd"
	"github.com/bd878/lesnotes_bot/chats/internal/server"
	"github.com/bd878/lesnotes_bot/chats/internal/logging"
	"github.com/bd878/lesnotes_bot/chats/internal/handlers"
	"github.com/bd878/lesnotes_bot/chats/internal/gateway"
	"github.com/bd878/lesnotes_bot/chats/internal/http"
	"github.com/bd878/lesnotes_bot/chats/internal/repository"
	"github.com/bd878/lesnotes_bot/chats/internal/application"
)

type Module struct {}

func (m *Module) Startup(ctx context.Context, mono system.Monolith) error {
	eventDispatcher := ddd.NewEventDispatcher[ddd.Event]()

	client, err := http.NewClient()
	if err != nil {
		return err
	}

	gallery := logging.LogGatewayHandlers(
		gateway.NewChatsGateway(client, mono.Config().UsersURL),
		mono.Log(),
	)

	domainHandlers := logging.LogDomainEventHandlers[ddd.Event](
		handlers.NewDomainHandlers(mono.Bot(), mono.Log()), mono.Log(),
	)

	chats := repository.NewChatsRepository("chats.chats", mono.Pool())
	app := logging.LogApplicationAccess(application.New(chats, gallery, eventDispatcher), mono.Log())

	if err := server.RegisterBot(app, mono.Bot(), mono.Log()); err != nil {
		return err
	}

	handlers.RegisterDomainEventHandlers(eventDispatcher, domainHandlers)

	return nil
}

func (m Module) Name() string {
	return "chats"
}