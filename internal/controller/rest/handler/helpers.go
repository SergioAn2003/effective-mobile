package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/SergioAn2003/effective-mobile/pkg/logger"
)

type ResponseError struct {
	Message string `json:"message"`
}

func sendErr(ctx context.Context, w http.ResponseWriter, code int, err error, msg string) {
	l := logger.FromContext(ctx)

	l.Error("api error", "error", err, "code", code)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	err = json.NewEncoder(w).Encode(ResponseError{Message: msg})
	if err != nil {
		l.Error("api error", "error", err, "code", http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type EmptyJSON struct{}

func sendJSON(ctx context.Context, w http.ResponseWriter, code int, data any) {
	if data == nil {
		data = EmptyJSON{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		sendErr(ctx, w, http.StatusInternalServerError, err, "Внутренняя ошибка")
		return
	}
}
