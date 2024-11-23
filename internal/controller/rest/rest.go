package rest

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/SergioAn2003/effective-mobile/internal/controller/rest/handler"
	"github.com/SergioAn2003/effective-mobile/internal/controller/rest/middlewares"
	"github.com/SergioAn2003/effective-mobile/pkg/config"
	"github.com/go-chi/chi/v5"
)

type Controller struct {
	cfg    config.Config
	log    *slog.Logger
	server *http.Server
}

func New(cfg config.Config, log *slog.Logger, service handler.Service) *Controller {
	router := chi.NewRouter()

	mw := middlewares.New(cfg, log)
	handler := handler.New(log, service)

	router.Use(mw.Log, mw.Recover, mw.Cors)

	router.Get("/ping", handler.Ping)
	router.Post("/songs", handler.CreateSong)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.HTTP.HTTPPort),
		Handler:      router,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
	}

	fmt.Println(server.Addr)

	return &Controller{
		cfg:    cfg,
		log:    log,
		server: server,
	}
}

func (r *Controller) Run() error {
	return r.server.ListenAndServe()
}

func (r *Controller) Shutdown(ctx context.Context) error {
	return r.server.Shutdown(ctx)
}
