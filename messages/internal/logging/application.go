package logging

import (
	"context"

	"go.uber.org/zap"

	"github.com/bd878/lesnotes_bot/internal/logger"
	"github.com/bd878/lesnotes_bot/messages/internal/application"
)

type Application struct {
	application.App
	logger *logger.Logger
}

var _ application.App = (*Application)(nil)

func LogApplicationAccess(application application.App, logger *logger.Logger) Application {
	return Application{
		App: application,
		logger: logger,
	}
}

func (a Application) CreateMessage(ctx context.Context, cmd application.CreateMessage) (err error) {
	a.logger.Infoln("--> CreateMessage")
	defer func() { a.logger.WithOptions(zap.Fields(zap.Error(err))).Infoln("<-- CreateMessage") }()
	return a.App.CreateMessage(ctx, cmd)
}
