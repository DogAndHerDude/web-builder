package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB 

func Init() {
  connection, err := sqlx.Connect("sqlite3", os.Getenv("DATABASE_URL"))

  if err != nil {
    log.Fatal(err)
  }

  DB = connection

  DB.MustExec(schema)
  log.Println("Schema synced")
}
