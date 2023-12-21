package storage

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/patrickmn/go-cache"
	"github.com/yawkar/wbl0/pkg/models"
)

type Storage struct {
	cache *cache.Cache
	db    *sqlx.DB
}

type StorageConfig struct {
	dbConfig
	cacheConfig
}

func MkStorage(c *StorageConfig, loadCaches bool) (*Storage, error) {
	db, err := mkDb(&c.dbConfig)
	if err != nil {
		return nil, errors.Join(errors.New("couldn't make database connection"), err)
	}
	cache, err := mkCache(&c.cacheConfig)
	if err != nil {
		return nil, errors.Join(errors.New("couldn't make cache"), err)
	}
	return &Storage{
		db:    db,
		cache: cache,
	}, nil
}

func (s *Storage) BeginTx() (*sql.Tx, error) {
	return s.db.Begin()
}

func (s *Storage) InsertOrder(order *models.Order) error {
	if err := insertOrder(s.db, order); err != nil {
		return err
	}
	cacheOrder(s.cache, order)
	return nil
}

func (s *Storage) InsertPayment(payment *models.Payment) error {
	if err := insertPayment(s.db, payment); err != nil {
		return err
	}
	cachePayment(s.cache, payment)
	return nil
}

func (s *Storage) InsertDelivery(delivery *models.Delivery) error {
	if err := insertDelivery(s.db, delivery); err != nil {
		return err
	}
	cacheDelivery(s.cache, delivery)
	return nil
}

func (s *Storage) InsertItem(item *models.Item) error {
	if err := insertItem(s.db, item); err != nil {
		return err
	}
	cacheItem(s.cache, item)
	return nil
}
