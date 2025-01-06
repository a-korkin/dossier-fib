package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func OpenDB(driver, conn string) error {
	var err error
	Db, err = sql.Open(driver, conn)
	return err
}

func CloseDB() error {
	return Db.Close()
}
