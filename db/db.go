package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
	"os"
)

func Connect() *sql.DB {

	var username string = os.Getenv("PG_USERNAME")
	var password string = os.Getenv("PG_PASSWORD")
	var database string = os.Getenv("PG_DATABASE")

	// Create connection to Postgres
	connInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, database)
	db, err := sql.Open("postgres", connInfo)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
