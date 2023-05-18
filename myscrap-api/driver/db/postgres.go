package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var postgresDB *sql.DB

type PostgresDB struct{}

func InitDB() {
	conn, err := sql.Open("postgres", "host=myscrap-db port=5432 user=myscrap-user password=myscrap-password dbname=myscrap sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	postgresDB = conn
	log.Println("DB INITIALIZED")
}

func NewPostgresDB() *sql.DB {
	return postgresDB
}

func FormatTimestamptz(t *time.Time) string {
	return t.Format(time.RFC3339)
}
