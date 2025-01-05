package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB(conn string) (*PostgresDB, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	pg := PostgresDB{
		DB: db,
	}

	return &pg, nil
}

func (pg *PostgresDB) Close() error {
	if err := pg.DB.Close(); err != nil {
		return err
	}
	return nil
}
