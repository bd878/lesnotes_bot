package logging

import (
	"context"

	"github.com/go-telegram/bot/models"
	botApi "github.com/go-telegram/bot"

	"github.com/bd878/lesnotes_bot/internal/logger"
	"github.com/bd878/lesnotes_bot/chats/internal/application"
)

type (
	App interface {
		Start(ctx context.Context, b *botApi.Bot, update *models.Update)
	}

	Application struct {
		application.App
		logger *logger.Logger
	}
)

var _ App = (*Application)(nil)

func LogApplicationAccess(application application.App, logger *logger.Logger) Application {
	return Application{
		App: application,
		logger: logger,
	}
}

func (a Application) Start(ctx context.Context, b *botApi.Bot, update *models.Update) {
	var err error
	a.logger.Infoln("---> Start")
	defer func() { a.logger.Infof("<-- Start: %s\n", err.Error()) }()
	err = a.App.Start(ctx, b, update)
}