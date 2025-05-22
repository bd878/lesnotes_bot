package main

import (
	"flag"
	"os"
	"fmt"
	"net/http"
	"context"

	"golang.org/x/sync/errgroup"
	"github.com/jackc/pgx/v5/pgxpool"
	botApi "github.com/go-telegram/bot"

	"github.com/bd878/lesnotes_bot/internal/logger"
	"github.com/bd878/lesnotes_bot/internal/waiter"
	"github.com/bd878/lesnotes_bot/internal/config"
	"github.com/bd878/lesnotes_bot/internal/bot"
	"github.com/bd878/lesnotes_bot/internal/system"
	"github.com/bd878/lesnotes_bot/chats"
	"github.com/bd878/lesnotes_bot/messages"
)

var help bool

func init() {
	flag.BoolVar(&help, "help", false, "show usage")

	flag.Usage = func() {
		fmt.Printf("Usage: %s path/to/config.json", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(); err != nil {
		logger.Log.Errorln(err)
	}

	logger.Log.Sync()
}

type app struct {
	cfg config.Config
	waiter waiter.Waiter
	log *logger.Logger
	pool *pgxpool.Pool
	bot *bot.Bot
	server *http.Server
	modules []system.Module
}

func (a *app) Pool() *pgxpool.Pool {
	return a.pool
}

func (a *app) Log() *logger.Logger {
	return a.log
}

func (a *app) Bot() *bot.Bot {
	return a.bot
}

func (a *app) Waiter() waiter.Waiter {
	return a.waiter
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) Modules() []system.Module {
	return a.modules
}

func (a *app) Server() *http.Server {
	return a.server
}

var _ system.Monolith = (*app)(nil)

func run() (err error) {
	a := &app{}

	a.cfg = config.LoadConfig(os.Args[1])
	a.log = logger.NewLog()
	a.bot = bot.New(
		os.Getenv("TELEGRAM_LESNOTES_BOT_TOKEN"),
		os.Getenv("TELEGRAM_LESNOTES_BOT_WEBHOOK_SECRET_TOKEN"),
		a.cfg.WebhookURL + a.cfg.WebhookPath,
		botApi.WithDebug(),
	)
	a.pool, err = pgxpool.New(context.Background(), a.cfg.PGConn)
	if err != nil {
		return err
	}

	a.modules = []system.Module{
		&chats.Module{},
		&messages.Module{},
	}

	a.waiter = waiter.New(waiter.CatchSignals())

	for _, module := range a.modules {
		if err = module.Startup(a.Waiter().Context(), a); err != nil {
			return err
		}
	}

	a.server = &http.Server{
		Addr: a.cfg.Addr,
	}
	mux := http.NewServeMux()
	mux.HandleFunc(a.cfg.WebhookPath, a.Bot().WebhookHandler())
	a.server.Handler = mux

	a.waiter.Add(
		a.waitForPool,
		a.waitForBot,
		a.waitForWeb,
	)

	return a.waiter.Wait()
}

func (a *app) waitForPool(ctx context.Context) error {
	group, gCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		<-gCtx.Done()
		a.log.Infoln("closing pgpool connections")
		a.pool.Close()
		return nil
	})

	return group.Wait()
}

func (a *app) waitForBot(ctx context.Context) error {
	group, gCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		a.log.Infow("start webhook", "url", a.cfg.WebhookURL + a.cfg.WebhookPath)
		defer a.log.Infoln("webhook shutdown")
		a.bot.StartWebhook(ctx)
		return nil
	})

	group.Go(func() (err error) {
		<-gCtx.Done()
		a.log.Infoln("webhook is about to be deleted")
		_, err = a.bot.DeleteWebhook(context.Background(), &botApi.DeleteWebhookParams{
			DropPendingUpdates: true,
		})
		a.log.Infoln("webhook deleted")
		return err
	})

	return group.Wait()
}

func (a *app) waitForWeb(ctx context.Context) error {
	group, gCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		a.log.Infow("start web server", "addr", a.cfg.Addr)
		err := a.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			a.log.Errorw("web server exited with error", "error", err)
		}
		return err
	})

	group.Go(func() (err error) {
		<-gCtx.Done()
		a.log.Infoln("web server is about to be shutdown")
		if err := a.Server().Shutdown(context.Background()); err != nil {
			a.log.Errorln("failed to stop web server")
		}
		return
	})

	return group.Wait()
}