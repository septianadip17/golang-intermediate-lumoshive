package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PgxIface interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

func InitDB() (*pgx.Conn, error) {
	connStr := "user=postgres password=postgres dbname=ojek_online sslmode=disable host=192.168.1.33"
	conn, err := pgx.Connect(context.Background(), connStr)
	return conn, err
}
