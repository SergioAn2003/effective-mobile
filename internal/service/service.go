package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/SergioAn2003/effective-mobile/internal/entity"
	"github.com/gofrs/uuid/v5"
)

type Service struct {
	postgresRepo postgresRepo
	songClient   songClient
}

func New(repo postgresRepo, songClient songClient) *Service {
	return &Service{
		postgresRepo: repo,
		songClient:   songClient,
	}
}

func (s *Service) CreateSong(ctx context.Context, songName, groupName string) (uuid.UUID, error) {
	_, err := s.postgresRepo.GetSongByNameAndGroupName(ctx, songName, groupName)
	if err == nil {
		return uuid.UUID{}, fmt.Errorf("method create song: get song by name and group name: %w", entity.ErrAlreadyExists)
	}

	if !errors.Is(err, entity.ErrNotFound) {
		return uuid.UUID{}, fmt.Errorf("method create song: %w", err)
	}

	songDetails, err := s.songClient.GetSongDetails(ctx, songName, groupName)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("method create song: get song details: %w", err)
	}

	newSong := entity.Song{
		ID:          uuid.Must(uuid.NewV4()),
		Name:        songName,
		GroupName:   groupName,
		SongDetails: songDetails,
	}

	if err := s.postgresRepo.CreateSong(ctx, newSong); err != nil {
		return uuid.UUID{}, fmt.Errorf("method create song: create song: %w", err)
	}

	return newSong.ID, nil
}
