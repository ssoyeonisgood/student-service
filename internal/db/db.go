package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	var err error

	host := "localhost"
	port := 5432
	user := "soyeon_go"
	password := "soyeon_go"
	dbname := "studentdb"

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	return DB.Ping()
}
