package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewConnection(dbUrl string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
