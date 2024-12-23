package db

import (
	"context"

	"github.com/hasanaburayyan/raid-bot/common/models"
	"github.com/jackc/pgx/v5"
)

func CreateRaider(conn *pgx.Conn, raider models.Raider) error {
	query := `INSERT INTO raiders (id, discord_id) VALUES ($1, $2)`
	_, err := conn.Exec(context.Background(), query, raider.ID, raider.DiscordID)
	return err
}

func GetRaiders(conn *pgx.Conn) ([]models.Raider, error) {
	query := `SELECT id, discord_id FROM raiders`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var raiders []models.Raider
	for rows.Next() {
		var raider models.Raider
		err := rows.Scan(&raider.ID, &raider.DiscordID)
		if err != nil {
			return nil, err
		}
		raiders = append(raiders, raider)
	}

	return raiders, nil
}
