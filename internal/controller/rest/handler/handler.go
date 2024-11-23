package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Handler struct {
	log     *slog.Logger
	service Service
}

func New(log *slog.Logger, service Service) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("pong"))
}

type CreateSongRequest struct {
	SongName  string `json:"song"`
	GroupName string `json:"group"`
}

type CreateSongResponse struct {
	Message string `json:"message"`
}

func (h *Handler) CreateSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req CreateSongRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendErr(ctx, w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	songID, err := h.service.CreateSong(ctx, req.SongName, req.GroupName)
	if err != nil {
		sendErr(ctx, w, http.StatusInternalServerError, err, "failed to create song")
		return
	}

	sendJSON(ctx, w, http.StatusCreated, CreateSongResponse{
		Message: "song successfully created, id: " + songID.String(),
	})
}
