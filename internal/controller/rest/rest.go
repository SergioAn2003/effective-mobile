package rest

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/SergioAn2003/effective-mobile/internal/controller/rest/handler"
	"github.com/SergioAn2003/effective-mobile/internal/controller/rest/router"
	"github.com/SergioAn2003/effective-mobile/pkg/config"
)

type Controller struct {
	cfg    config.Config
	log    *slog.Logger
	server *http.Server
}

func New(cfg config.Config, log *slog.Logger, service handler.Service) *Controller {
	router := router.New(cfg, log, service)

	server := &http.Server{
		Addr:         ":" + cfg.HTTP.HTTPPort,
		Handler:      router,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
	}

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
