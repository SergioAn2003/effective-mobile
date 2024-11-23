package postgres

import (
	"context"
	"errors"

	"github.com/SergioAn2003/effective-mobile/internal/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) GetSongByNameAndGroupName(ctx context.Context, songName, groupName string) (entity.Song, error) {
	sqlQuery := `
		SELECT id, song_name, group_name, release_date, lyrics, link
		FROM songs
		WHERE song_name = $1 AND group_name = $2`

	var song entity.Song

	err := r.pool.QueryRow(ctx, sqlQuery, songName, groupName).Scan(
		&song.ID,
		&song.Name,
		&song.GroupName,
		&song.SongDetails.ReleaseDate,
		&song.SongDetails.Lyrics,
		&song.SongDetails.Link,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Song{}, entity.ErrNotFound
		}

		return entity.Song{}, err
	}

	return song, nil
}

func (r *Repository) CreateSong(ctx context.Context, song entity.Song) error {
	sqlQuery := `
		INSERT INTO songs (id, song_name, group_name, release_date, lyrics, link)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.pool.Exec(ctx, sqlQuery,
		song.ID,
		song.Name,
		song.GroupName,
		song.SongDetails.ReleaseDate,
		song.SongDetails.Lyrics,
		song.SongDetails.Link,
	)

	if err != nil {
		return err
	}

	return nil
}
