package messages

import (
	"context"

	"github.com/bd878/lesnotes_bot/internal/system"
	"github.com/bd878/lesnotes_bot/internal/ddd"
	"github.com/bd878/lesnotes_bot/messages/internal/server"
	"github.com/bd878/lesnotes_bot/messages/internal/logging"
	"github.com/bd878/lesnotes_bot/messages/internal/handlers"
	"github.com/bd878/lesnotes_bot/messages/internal/http"
	"github.com/bd878/lesnotes_bot/messages/internal/grpc"
	"github.com/bd878/lesnotes_bot/messages/internal/gateway"
	"github.com/bd878/lesnotes_bot/messages/internal/application"
)

type Module struct {}

func (m *Module) Startup(ctx context.Context, mono system.Monolith) error {
	// driven
	eventDispatcher := ddd.NewEventDispatcher[ddd.Event]()

	httpClient, err := http.NewClient()
	if err != nil {
		return err
	}

	grpcClient, err := grpc.NewClient(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}

	chats := grpc.NewChatRepository(grpcClient)
	messages := gateway.NewMessagesGateway(httpClient, mono.Config().MessagesURL)

	domainHandlers := logging.LogDomainEventHandlers[ddd.Event](
		handlers.NewDomainHandlers(mono.Bot(), mono.Log()), mono.Log(),
	)

	// application
	app := logging.LogApplicationAccess(application.New(messages, chats, eventDispatcher), mono.Log())

	// driver
	if err := server.RegisterBot(app, mono.Bot(), mono.Log()); err != nil {
		return err
	}

	handlers.RegisterDomainEventHandlers(eventDispatcher, domainHandlers)

	return nil
}

func (m Module) Name() string {
	return "messages"
}