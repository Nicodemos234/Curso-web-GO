package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	connection := "user=postgres dbname=nico_loja password=123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
