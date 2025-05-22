package messages

import (
	"context"

	"github.com/bd878/lesnotes_bot/internal/system"
	"github.com/bd878/lesnotes_bot/messages/internal/server"
	"github.com/bd878/lesnotes_bot/messages/internal/logging"
	"github.com/bd878/lesnotes_bot/messages/internal/http"
	"github.com/bd878/lesnotes_bot/messages/internal/gateway"
	"github.com/bd878/lesnotes_bot/messages/internal/application"
)

type Module struct {}

func (m *Module) Startup(ctx context.Context, mono system.Monolith) error {
	client, err := http.NewClient()
	if err != nil {
		return err
	}

	messages := gateway.NewMessagesGateway(client, mono.Config().MessagesURL)
	app := logging.LogApplicationAccess(application.New(messages), mono.Log())

	if err := server.RegisterBot(app, mono.Bot(), mono.Log()); err != nil {
		return err
	}

	return nil
}

func (m Module) Name() string {
	return "messages"
}