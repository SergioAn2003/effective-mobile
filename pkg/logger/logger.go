package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/SergioAn2003/effective-mobile/internal/entity"
)

const (
	JSON = "JSON"
	TEXT = "TEXT"
)

type handler struct {
	slog.Handler
}

func New(level, format string) (*slog.Logger, error) {
	var (
		h         *handler
		slogLevel slog.Level
	)

	if err := slogLevel.UnmarshalText([]byte(level)); err != nil {
		return nil, err
	}

	opts := &slog.HandlerOptions{Level: slogLevel}

	switch strings.ToUpper(format) {
	case JSON:
		h = &handler{slog.NewJSONHandler(os.Stdout, opts)}
	case TEXT:
		h = &handler{slog.NewTextHandler(os.Stdout, opts)}
	default:
		return nil, fmt.Errorf("invalid logger format: %s", format)
	}

	logger := slog.New(h)

	return logger, nil
}

func FromContext(ctx context.Context) *slog.Logger {
	log, ok := ctx.Value(entity.CtxKeyLogger{}).(*slog.Logger)
	if !ok {
		return slog.New(slog.NewJSONHandler(io.Discard, nil))
	}

	return log
}
