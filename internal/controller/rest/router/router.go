package router

import (
	"log/slog"
	"net/http"

	"github.com/SergioAn2003/effective-mobile/internal/controller/rest/handler"
	"github.com/SergioAn2003/effective-mobile/internal/controller/rest/middlewares"
	"github.com/SergioAn2003/effective-mobile/pkg/config"
	"github.com/go-chi/chi/v5"
)

func New(cfg config.Config, log *slog.Logger, service handler.Service) http.Handler {
	router := chi.NewRouter()

	mw := middlewares.New(cfg, log)
	h := handler.New(log, service)

	router.Use(mw.Log, mw.Recover, mw.Cors)

	router.Route("/api", func(r chi.Router) {
		r.Get("/ping", h.Ping)
		r.Post("/songs", h.CreateSong)
		r.Delete("/songs/{id}", h.DeleteSong)
	})

	return router
}
