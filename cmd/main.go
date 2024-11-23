package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/SergioAn2003/effective-mobile/internal/controller/rest"
	"github.com/SergioAn2003/effective-mobile/internal/repository/postgres"
	"github.com/SergioAn2003/effective-mobile/internal/repository/songclient"
	"github.com/SergioAn2003/effective-mobile/internal/service"
	"github.com/SergioAn2003/effective-mobile/pkg/config"
	"github.com/SergioAn2003/effective-mobile/pkg/logger"
	"github.com/SergioAn2003/effective-mobile/pkg/pg"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.New(".env")
	panicOnErr("failed to load config", err)

	log, err := logger.New(cfg.Logger.Level, cfg.Logger.Format)
	panicOnErr("failed to create logger", err)

	pool, err := pg.Connect(ctx, cfg.Postgres.DSN, cfg.Postgres.MaxConns)
	panicOnErr("failed to connect to database", err)

	err = pg.UpMigrations(ctx, pool)
	panicOnErr("failed to run migrations", err)

	postgresRepo := postgres.New(pool)
	songClient := songclient.New(cfg.ExternalSongService.BaseURL)

	service := service.New(postgresRepo, songClient)
	restController := rest.New(cfg, log, service)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		err = restController.Run()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("failed to run rest server", "error", err)
			return
		}
	}()

	log.Info("rest server started", "port", cfg.HTTP.HTTPPort)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	sig := <-ch

	log.Info("received signal", "signal", sig.String())

	err = restController.Shutdown(ctx)
	if err != nil {
		log.Error("failed to shutdown rest server", "error", err)
	}

	wg.Wait()
}

func panicOnErr(msg string, err error) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
