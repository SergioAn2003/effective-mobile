package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTP                HTTP
	Logger              Logger
	Postgres            Postgres
	ExternalSongService ExternalSongService
}

type HTTP struct {
	HTTPPort     string        `env:"HTTP_PORT"`
	ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT"`
	WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT"`
}

type Logger struct {
	Level  string `env:"LOGGER_LEVEL"`
	Format string `env:"LOGGER_FORMAT"`
}

type Postgres struct {
	DSN      string `env:"POSTGRES_DSN"`
	MaxConns int32  `env:"POSTGRES_MAX_CONNS"`
}

type ExternalSongService struct {
	BaseURL string `env:"SONG_SERVICE_BASE_URL"`
}

func New(envPath string) (Config, error) {
	err := godotenv.Load(envPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return Config{}, fmt.Errorf("failed to load config: %w", err)
	}

	c, err := env.ParseAsWithOptions[Config](env.Options{RequiredIfNoDef: true})
	if err != nil {
		return Config{}, fmt.Errorf("failed to parse config: %w", err)
	}

	return c, nil
}
