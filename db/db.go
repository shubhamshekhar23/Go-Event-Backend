package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Database connection failed: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = CreateAllTables()
	if err != nil {
		panic("Failed to create tables: " + err.Error())
	}

	return nil
}
