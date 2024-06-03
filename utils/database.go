package utils

import (
	"client-server-challenge-go/config"
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

func InitDatabase() (*sql.DB, error) {
	if _, err := os.Stat(config.DataFolder); os.IsNotExist(err) {
		err := os.Mkdir(config.DataFolder, 0755)
		if err != nil {
			return nil, err
		}
	}

	db, err := sql.Open("sqlite3", config.DataFolder+config.DbFilePath)
	if err != nil {
		log.Fatalln("Error opening database:", err)
		return nil, err
	}
	log.Println("Database file created:", config.DataFolder+config.DbFilePath)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS bids (id INTEGER PRIMARY KEY AUTOINCREMENT, bid DOUBLE, created_at DATETIME)`)
	if err != nil {
		log.Fatalln("Error creating table:", err)
		return nil, err
	}

	return db, nil
}

func SaveData(db *sql.DB, bid float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), config.DbTimeout)
	defer cancel()

	_, err := db.ExecContext(ctx, "INSERT INTO bids (bid, created_at) VALUES (?, ?)", bid, time.Now())
	return err
}
