package db

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jackc/pgx/v5"
)

func RunMigrations(conn *pgx.Conn, migrationsDir string) error {
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".up.sql") {
			continue
		}

		// Read SQL file contents
		path := filepath.Join(migrationsDir, file.Name())
		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		_, err = conn.Exec(context.Background(), string(contents))
		if err != nil {
			return err
		}

		log.Printf("Applied migration: %s", file.Name())
	}

	return nil
}
