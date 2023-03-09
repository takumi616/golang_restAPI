package db

import (
	"database/sql"
	"log"
	"os"

    _ "github.com/lib/pq"
)

func Init() *sql.DB {
	dbUser := os.Getenv("PGDB_USER")
	dbPass := os.Getenv("PGDB_PASSWORD")
	dbName := os.Getenv("PGDB_NAME_TASK_MANAGEMENT_APP")
	port := os.Getenv("PGDB_PORT")
	sslmode := os.Getenv("PGDB_SSLMODE")

	dataSourceName := "port=" + port + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " sslmode=" + sslmode
	db, err := sql.Open("postgres", dataSourceName)
    // defer db.Close()

    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    } else {
		log.Println("Connecting to database is success.")
	}

	return db
}