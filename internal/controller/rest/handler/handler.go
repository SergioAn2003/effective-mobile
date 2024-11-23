package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/SergioAn2003/effective-mobile/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
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

func (h *Handler) Ping(w http.ResponseWriter, _ *http.Request) {
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
		if errors.Is(err, entity.ErrAlreadyExists) {
			sendErr(ctx, w, http.StatusConflict, err, "song already exists")
			return
		}

		sendErr(ctx, w, http.StatusInternalServerError, err, "failed to create song")
		return
	}

	sendJSON(ctx, w, http.StatusCreated, CreateSongResponse{
		Message: "song successfully created, id: " + songID.String(),
	})
}

type DeleteSongResponse struct {
	Message string `json:"message"`
}

func (h *Handler) DeleteSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	qSongID := chi.URLParam(r, "id")

	songID, err := uuid.FromString(qSongID)
	if err != nil {
		sendErr(ctx, w, http.StatusBadRequest, err, "invalid song id: "+qSongID)
		return
	}

	err = h.service.DeleteSong(ctx, songID)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			sendErr(ctx, w, http.StatusNotFound, err, "song not found")
			return
		}

		sendErr(ctx, w, http.StatusInternalServerError, err, "failed to delete song")
		return
	}

	sendJSON(ctx, w, http.StatusOK, DeleteSongResponse{
		Message: "song successfully deleted",
	})
}
