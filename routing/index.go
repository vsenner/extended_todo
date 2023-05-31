package routing

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func SetupDB() {
	dbConnStr := "postgresql://postgres:a86UXJwyCIMpjzSdlVvL@containers-us-west-93.railway.app:6097/railway"

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
