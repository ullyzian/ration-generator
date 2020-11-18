package db

import (
	"database/sql"
	"log"
)

func CreateTable(conn *sql.DB) {
	cmd := `CREATE TABLE IF NOT EXISTS 
			dishes(id SERIAL PRIMARY KEY, title VARCHAR(128) UNIQUE NOT NULL, portion INT NOT NULL, calories INT NOT NULL)`
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}
}