package repository

import (
    "context"

    "ainyx/config"
    "ainyx/db/sqlc"

    "github.com/jackc/pgx/v5/pgxpool"
)


func ConnectDB(cfg *config.Config) (*sqlc.Queries, *pgxpool.Pool, error) {

    pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
    if err != nil {
        return nil, nil, err
    }

    db := sqlc.New(pool)
    return db, pool, nil
}
