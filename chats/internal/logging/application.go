package logging

import (
	"context"

	"go.uber.org/zap"

	"github.com/bd878/lesnotes_bot/internal/logger"
	"github.com/bd878/lesnotes_bot/chats/internal/application"
	"github.com/bd878/lesnotes_bot/chats/internal/domain"
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

func (a Application) CreateChat(ctx context.Context, cmd application.CreateChat) (err error) {
	a.logger.Infoln("--> CreateChat")
	defer func() { a.logger.WithOptions(zap.Fields(zap.Error(err))).Infoln("<-- CreateChat") }()
	return a.App.CreateChat(ctx, cmd)
}

func (a Application) GetChat(ctx context.Context, cmd application.GetChat) (chat *domain.Chat, err error) {
	a.logger.Infoln("--> GetChat")
	defer func() { a.logger.WithOptions(zap.Fields(zap.Error(err))).Infoln("<-- GetChat") }()
	return a.App.GetChat(ctx, cmd)
}

func (a Application) KickMember(ctx context.Context, cmd application.KickMember) (err error) {
	a.logger.Infoln("--> KickMember")
	defer func() { a.logger.WithOptions(zap.Fields(zap.Error(err))).Infoln("<-- KickMember") }()
	return a.App.KickMember(ctx, cmd)
}
