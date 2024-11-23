package entity

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Song struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"song_name"`
	GroupName   string      `json:"group_name"`
	SongDetails SongDetails `json:"song_details"`
}

type SongDetails struct {
	ReleaseDate time.Time `json:"release_date"`
	Lyrics      string    `json:"lyrics"`
	Link        string    `json:"link"`
}
