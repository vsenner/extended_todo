package routing

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func SetupDB() {
	dbConnStr := "postgres://postgres:postgres@localhost:5432/extended_todo?sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
