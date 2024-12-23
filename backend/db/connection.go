package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

func DefaultConnection() (*pgx.Conn, error) {
	dbUrl := os.Getenv("DATABASE_URL")
	return NewConnection(dbUrl)
}

func NewConnection(dbUrl string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
