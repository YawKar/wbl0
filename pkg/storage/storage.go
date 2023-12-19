package storage

import (
	"database/sql"

	"github.com/patrickmn/go-cache"
	"github.com/yawkar/wbl0/pkg/models"
)

type Storage struct {
	cache *cache.Cache
	db    *sql.DB
}

type StorageConfig struct {
	dbConfig
	cacheConfig
}

func MkStorage(c *StorageConfig, loadCaches bool) (*Storage, error) {
	// TODO
}

func (s *Storage) BeginTx() (*sql.Tx, error) {
	return s.db.Begin()
}

func (s *Storage) InsertOrder(order *models.Order) error {
	tx, err := s.db.Begin()
	if err != nil {
		return nil
	}
	defer tx.Rollback()
	// TODO
}

func (s *Storage) InsertPayment(payment *models.Payment) error {
	tx, err := s.db.Begin()
	if err != nil {
		return nil
	}
	defer tx.Rollback()
	// TODO
}

func (s *Storage) InsertDelivery(delivery *models.Delivery) error {
	tx, err := s.db.Begin()
	if err != nil {
		return nil
	}
	defer tx.Rollback()
	// TODO
}

func (s *Storage) InsertItem(item *models.Item) error {
	tx, err := s.db.Begin()
	if err != nil {
		return nil
	}
	defer tx.Rollback()
	// TODO
}
