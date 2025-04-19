package main

import (
	"flag"
	"os"
	"fmt"
	"context"

	"golang.org/x/sync/errgroup"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/bd878/lesnotes_bot/internal/logger"
	"github.com/bd878/lesnotes_bot/internal/waiter"
	"github.com/bd878/lesnotes_bot/internal/config"
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
}

func run() error {
	a := app{}

	a.cfg = config.LoadConfig(os.Args[1])
	a.log = logger.NewLog()

	a.waiter = waiter.New()
	a.waiter.Add(
		a.waitForPool,
		a.waitForWeb,
	)

	return a.waiter.Wait()
}

func (a *app) waitForPool(ctx context.Context) error {
	group, gCtx := errgroup.WithContext(ctx)

	group.Go(func() (err error) {
		a.log.Infoln("start pgpool")
		defer a.log.Infoln("pgpool shutdown")
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

func (a *app) waitForWeb(ctx context.Context) error {
	return nil
}