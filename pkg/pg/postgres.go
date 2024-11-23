package pg

import (
	"context"
	"errors"
	"fmt"

	"github.com/SergioAn2003/effective-mobile/migrations"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func Connect(ctx context.Context, dsn string, maxConns int32) (*pgxpool.Pool, error) {
	pgxConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to parse config: %w", err)
	}

	pgxConfig.MaxConns = maxConns

	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to connect to database: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("postgres: failed to ping database: %w", err)
	}

	return pool, nil
}

func UpMigrations(ctx context.Context, pool *pgxpool.Pool) error {
	db := stdlib.OpenDBFromPool(pool)

	fs := migrations.FS
	goose.SetBaseFS(fs)
	goose.SetLogger(goose.NopLogger())

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("postgres: failed to set dialect: %w", err)
	}

	err := goose.Up(db, ".")
	if err != nil && !errors.Is(err, goose.ErrNoNextVersion) {
		return fmt.Errorf("postgres: failed to run migrations: %w", err)
	}

	return nil
}
