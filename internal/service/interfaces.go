package service

import (
	"context"

	"github.com/SergioAn2003/effective-mobile/internal/entity"
)

type postgresRepo interface {
	GetSongByNameAndGroupName(ctx context.Context, songName, groupName string) (entity.Song, error)
	CreateSong(ctx context.Context, song entity.Song) error
}

type songClient interface {
	GetSongDetails(ctx context.Context, songName, groupName string) (entity.SongDetails, error)
}
