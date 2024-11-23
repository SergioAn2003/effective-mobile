package config_test

import (
	"testing"
	"time"

	"github.com/SergioAn2003/effective-mobile/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	r := require.New(t)

	t.Setenv("HTTP_PORT", "8080")
	t.Setenv("HTTP_READ_TIMEOUT", "10s")
	t.Setenv("HTTP_WRITE_TIMEOUT", "10s")
	t.Setenv("LOGGER_LEVEL", "info")
	t.Setenv("LOGGER_FORMAT", "json")
	t.Setenv("POSTGRES_DSN", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	t.Setenv("POSTGRES_MAX_CONNS", "5")

	expectedConfig := config.Config{
		HTTP: config.HTTP{
			HTTPPort:     "8080",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		Logger: config.Logger{
			Level:  "info",
			Format: "json",
		},
		Postgres: config.Postgres{
			DSN:      "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
			MaxConns: 5,
		},
	}

	cfg, err := config.New(".env.test")
	r.NoError(err)
	r.Equal(expectedConfig, cfg)
}
