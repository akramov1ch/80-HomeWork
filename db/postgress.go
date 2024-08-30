package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := "user=postgres password=vakhaboff dbname=imtihon_5 sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s\n", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %s\n", err)
	}

	fmt.Println("Connected to PostgreSQL database!")
}
