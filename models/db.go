package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "db.sqlite"

func ConnectToDB() *sql.DB {
	db, err := sql.Open("sqlite3", file)

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	return db
}
