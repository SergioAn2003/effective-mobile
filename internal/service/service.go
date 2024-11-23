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
		return uuid.UUID{}, fmt.Errorf("service.CreateSong(): get song by name and group name: %w", entity.ErrAlreadyExists)
	}

	if !errors.Is(err, entity.ErrNotFound) {
		return uuid.UUID{}, fmt.Errorf("service.CreateSong(): get song by name and group name: %w", err)
	}

	songDetails, err := s.songClient.GetSongDetails(ctx, songName, groupName)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("service.CreateSong(): get song details: %w", err)
	}

	newSong := entity.Song{
		ID:          uuid.Must(uuid.NewV4()),
		Name:        songName,
		GroupName:   groupName,
		SongDetails: songDetails,
	}

	if err := s.postgresRepo.CreateSong(ctx, newSong); err != nil {
		return uuid.UUID{}, fmt.Errorf("service.CreateSong(): create song: %w", err)
	}

	return newSong.ID, nil
}

func (s *Service) DeleteSong(ctx context.Context, songID uuid.UUID) error {
	_, err := s.postgresRepo.GetSongByID(ctx, songID)
	if err != nil {
		return fmt.Errorf("service.DeleteSong(): get song by id: %w", err)
	}

	if err := s.postgresRepo.DeleteSong(ctx, songID); err != nil {
		return fmt.Errorf("service.DeleteSong(): delete song: %w", err)
	}

	return nil
}
