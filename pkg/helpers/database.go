package helpers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./shortener.db")
	if err != nil {
		return db, err
	}
	return db, nil
}
