package handler

import (
	"context"

	"github.com/gofrs/uuid/v5"
)

type Service interface {
	CreateSong(ctx context.Context, songName, groupName string) (uuid.UUID, error)
	DeleteSong(ctx context.Context, songID uuid.UUID) error
}
