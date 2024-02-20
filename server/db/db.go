package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func New() *sqlx.DB {
	connection, err := sqlx.Connect("sqlite3", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	return connection
}
