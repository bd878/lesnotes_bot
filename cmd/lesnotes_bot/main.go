package main

import (
	"flag"
	"os"
	"fmt"
	"context"

	"golang.org/x/sync/errgroup"
	"github.com/jackc/pgx/v5/pgxpool"
	botApi "github.com/go-telegram/bot"

	"github.com/bd878/lesnotes_bot/internal/logger"
	"github.com/bd878/lesnotes_bot/internal/waiter"
	"github.com/bd878/lesnotes_bot/internal/config"
	"github.com/bd878/lesnotes_bot/internal/bot"
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
}

func run() error {
	a := app{}

	a.cfg = config.LoadConfig(os.Args[1])
	a.log = logger.NewLog()
	a.bot = bot.New(
		os.Getenv("TELEGRAM_LESNOTES_BOT_TOKEN"),
		os.Getenv("TELEGRAM_LESNOTES_BOT_WEBHOOK_SECRET_TOKEN"),
		a.cfg.WebhookURL + a.cfg.WebhookPath,
		botApi.WithDebug(),
	)

	a.waiter = waiter.New()
	a.waiter.Add(
		a.waitForPool,
		a.waitForBot,
	)

	return a.waiter.Wait()
}

func (a *app) waitForPool(ctx context.Context) error {
	group, gCtx := errgroup.WithContext(ctx)

	group.Go(func() (err error) {
		a.log.Infow("start pgpool", "conn", a.cfg.PGConn)
		if a.pool, err = pgxpool.New(context.Background(), a.cfg.PGConn); err != nil {
			a.log.Errorln("failed to create pgpool, exit")
		}
		return
	})

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