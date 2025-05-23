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

var _ domain.ChatGateway = (*gatewayHandlers)(nil)

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

func (g gatewayHandlers) Auth(ctx context.Context, token string) (err error) {
	g.logger.Infoln("--> Auth")
	defer func() { g.logger.WithOptions(zap.Fields(zap.Error(err))).Infoln("<-- Auth") }()
	return g.ChatGateway.Auth(ctx, token)
}

func (g gatewayHandlers) Login(ctx context.Context, login, password string) (token string, err error) {
	g.logger.Infoln("--> Login")
	defer func() { g.logger.WithOptions(zap.Fields(zap.Error(err))).Infoln("<-- Login") }()
	return g.ChatGateway.Login(ctx, login, password)
}

func (g gatewayHandlers) Delete(ctx context.Context, login, password, token string) (err error) {
	g.logger.Infoln("--> Delete")
	defer func() { g.logger.WithOptions(zap.Fields(zap.Error(err))).Infoln("<-- Delete") }()
	return g.ChatGateway.Delete(ctx, login, password, token)
}