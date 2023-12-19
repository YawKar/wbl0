package storage

import (
	"database/sql"
	"log"
	"time"

	"github.com/lib/pq"
	"github.com/patrickmn/go-cache"
)

type DbConfig struct {
	DbUrl string
}

func MkDb(c *DbConfig) (*sql.DB, error) {
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

type CacheConfig struct {
	CacheExpiration time.Duration
	CleanupInterval time.Duration
}

func MkCache(c *CacheConfig) (*cache.Cache, error) {
	cache := cache.New(c.CacheExpiration, c.CleanupInterval)
	return cache, nil
}
