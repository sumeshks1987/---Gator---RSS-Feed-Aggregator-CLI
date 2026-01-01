package main

import (
	"context"

	"rss/internal/database"

	"github.com/jackc/pgx/v5/pgxpool"
)

type State struct {
	DB     *database.Queries
	Config *Config
}

func NewState() (*State, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.New(context.Background(), cfg.DBUrl)
	if err != nil {
		return nil, err
	}

	return &State{
		DB:     database.New(pool),
		Config: cfg,
	}, nil
}
