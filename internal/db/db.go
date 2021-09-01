package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ozonva/ova-game-api/internal/configs"
)

func Connect(ctx context.Context, config *configs.Database) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?pool_max_conns=%d",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
		config.PoolMaxConnect,
	)
	pgxConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(ctx, pgxConfig)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
