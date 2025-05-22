package logging

import (
	"context"
	"go.uber.org/zap"

	"github.com/bd878/lesnotes_bot/internal/logger"
	"github.com/bd878/lesnotes_bot/chats/internal/domain"
)

type gatewayHandlers struct {
	domain.ChatGateway
	logger *logger.Logger
}

func LogGatewayHandlers(gateway domain.ChatGateway, logger *logger.Logger) gatewayHandlers {
	return gatewayHandlers{
		ChatGateway: gateway,
		logger: logger,
	}
}

func (g gatewayHandlers) Signup(ctx context.Context, login, password string) (token string, err error) {
	g.logger.Infoln("--> Signup")
	defer func() { g.logger.WithOptions(zap.Fields(zap.Error(err))).Infoln("<-- Signup") }()
	return g.ChatGateway.Signup(ctx, login, password)
}

func (g gatewayHandlers) Login(ctx context.Context, login, password string) (token string, err error) {
	g.logger.Infoln("--> Login")
	defer func() { g.logger.WithOptions(zap.Fields(zap.Error(err))).Infoln("<-- Login") }()
	return g.ChatGateway.Login(ctx, login, password)
}