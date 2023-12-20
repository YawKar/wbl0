package storage

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type dbConfig struct {
	DbUrl string
}

func mkDb(c *dbConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", c.DbUrl)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Fatalln("Couldn't ping db:", err)
	}
	return db, nil
}
