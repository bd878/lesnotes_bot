package system

import (
	"context"
	"google.golang.org/grpc"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/bd878/lesnotes_bot/internal/config"
	"github.com/bd878/lesnotes_bot/internal/logger"
	"github.com/bd878/lesnotes_bot/internal/bot"
)

type Monolith interface {
	Pool() *pgxpool.Pool
	Bot() *bot.Bot
	Log() *logger.Logger
	Config() config.Config
	Modules() []Module
	RPC() *grpc.Server
}

type Module interface {
	Startup(ctx context.Context, mono Monolith) error
	Name() string
}