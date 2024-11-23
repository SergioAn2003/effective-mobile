package service

import (
	"context"

	"github.com/SergioAn2003/effective-mobile/internal/entity"
	"github.com/gofrs/uuid/v5"
)

type postgresRepo interface {
	GetSongByID(ctx context.Context, songID uuid.UUID) (entity.Song, error)
	GetSongByNameAndGroupName(ctx context.Context, songName, groupName string) (entity.Song, error)
	CreateSong(ctx context.Context, song entity.Song) error
	DeleteSong(ctx context.Context, songID uuid.UUID) error
}

type songClient interface {
	GetSongDetails(ctx context.Context, songName, groupName string) (entity.SongDetails, error)
}
