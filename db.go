package main

import (
	"database/sql"
	"log"
)

var db *sql.DB

func GetConnection() *sql.DB {
	if db != nil {
		return db
	}

	db, err := sql.Open("sqlite3", "notesDB.sqlite")
	if err != nil {
		log.Fatalf("🔥 failed to connect to the database: %s", err.Error())
	}

	log.Println("🚀 Connected Successfully to the Database")

	return db
}

func MakeMigrations() error {
	db := GetConnection()

	stmt := `CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(64) UNIQUE CHECK(title IS NULL OR length(title) <= 64),
		description VARCHAR(255) NULL,
		completed BOOLEAN DEFAULT(FALSE),
		created_at TIMESTAMP DEFAULT DATETIME
	  );`

	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}
