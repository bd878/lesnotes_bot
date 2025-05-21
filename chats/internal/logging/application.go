package logging

import (
	"context"

	"go.uber.org/zap"

	"github.com/bd878/lesnotes_bot/internal/logger"
	"github.com/bd878/lesnotes_bot/chats/internal/application"
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
	a.logger.Infoln("---> CreateChat")
	defer func() { a.logger.WithOptions(zap.Fields(zap.Error(err))).Infoln("<-- CreateChat") }()
	return a.App.CreateChat(ctx, cmd)
}

func (a Application) CreateMessage(ctx context.Context, cmd application.CreateMessage) (res int32, err error) {
	a.logger.Infoln("---> CreateMessage")
	defer func() { a.logger.WithOptions(zap.Fields(zap.Error(err))).Infoln("<-- CreateMessage") }()
	return a.App.CreateMessage(ctx, cmd)
}

func (a Application) KickMember(ctx context.Context, cmd application.KickMember) (err error) {
	a.logger.Infoln("---> KickMember")
	defer func() { a.logger.WithOptions(zap.Fields(zap.Error(err))).Infoln("<-- KickMember") }()
	return a.App.KickMember(ctx, cmd)
}

func (a Application) ConfirmIssue(ctx context.Context, cmd application.ConfirmIssue) (err error) {
	a.logger.Infoln("---> ConfirmIssue")
	defer func() { a.logger.WithOptions(zap.Fields(zap.Error(err))).Infoln("<-- ConfirmIssue") }()
	return a.App.ConfirmIssue(ctx, cmd)
}