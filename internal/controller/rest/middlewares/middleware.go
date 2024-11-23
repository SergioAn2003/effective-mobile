package middlewares

import (
	"context"
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/SergioAn2003/effective-mobile/internal/entity"
	"github.com/SergioAn2003/effective-mobile/pkg/config"
	"github.com/gofrs/uuid/v5"
)

type Middleware struct {
	cfg config.Config
	log *slog.Logger
}

func New(cfg config.Config, log *slog.Logger) *Middleware {
	return &Middleware{
		cfg: cfg,
		log: log,
	}
}

func (m *Middleware) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := m.log.With("request_id", uuid.Must(uuid.NewV4()))

		l.Info("incoming request", "method", r.Method, "url", r.URL.String(), "from", r.RemoteAddr)

		ctx := context.WithValue(r.Context(), entity.CtxKeyLogger{}, l)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				m.log.Error("panic", "error", err, "stack", string(debug.Stack()))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin, Accept, User-Agent, Cache-Control")

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}
