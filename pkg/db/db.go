package db

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)


func InitDB() (*sql.DB, error) {

  db, err := sql.Open("sqlite3", "worktion.db")
  if err != nil {
    return nil, err
  }

  return db, nil
}
