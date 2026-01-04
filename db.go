package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB(filepath string) {
	var err error
	db, err = sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	//statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, passkey TEXT UNIQUE)")
	//statement.Exec()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, passkey TEXT UNIQUE)")
	if err != nil {
		log.Fatal(err)
	}
}

func isAllowedPasskey(passkey string) bool {
	var exists bool

	query := "SELECT EXISTS(SELECT 1 FROM users WHERE passkey=?)"
	err := db.QueryRow(query, passkey).Scan(&exists)
	if err != nil {
		log.Println("Database check error:", err)
		return false
	}
	return exists
}
