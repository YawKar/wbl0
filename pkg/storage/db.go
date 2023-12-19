package storage

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
)

type dbConfig struct {
	DbUrl string
}

func mkDb(c *dbConfig) (*sql.DB, error) {
	pqc, err := pq.NewConnector(c.DbUrl)
	if err != nil {
		return nil, err
	}
	db := sql.OpenDB(pqc)
	if err = db.Ping(); err != nil {
		log.Fatalln("Couldn't ping db:", err)
	}
	return db, nil
}
