package storage

import (
	"context"
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

var insertOrderSQL string = `
INSERT INTO "order" VALUES (
	:order_uid,
	:track_number,
	:entry,
	:locale,
	:internal_signature,
	:customer_id,
	:delivery_service,
	:shardkey,
	:sm_id,
	:date_created,
	:oof_shard
);`

func (s *Storage) InsertOrder(order *models.Order) error {
	tx, err := s.db.BeginTxx(context.TODO(), nil)
	if err != nil {
		return nil
	}
	defer tx.Rollback()
	if _, err := tx.NamedExec(insertOrderSQL, order); err != nil {
		return err
	}
	return tx.Commit()
}

var insertPaymentSQL string = `
INSERT INTO payment (
	transaction,
	request_id,
	currency,
	provider,
	amount,
	payment_dt,
	bank,
	delivery_cost,
	goods_total,
	custom_fee
) VALUES (
	:transaction,
	:request_id,
	:currency,
	:provider,
	:amount,
	:payment_dt,
	:bank,
	:delivery_cost,
	:goods_total,
	:custom_fee
);`

func (s *Storage) InsertPayment(payment *models.Payment) error {
	tx, err := s.db.BeginTxx(context.TODO(), nil)
	if err != nil {
		return nil
	}
	defer tx.Rollback()
	if _, err := tx.NamedExec(insertPaymentSQL, payment); err != nil {
		return err
	}
	return tx.Commit()
}

var insertDeliverySQL string = `
INSERT INTO delivery (
	order_uid,
	name,
	phone,
	zip,
	city,
	address,
	region,
	email
) VALUES (
	:order_uid,
	:name,
	:phone,
	:zip,
	:city,
	:address,
	:region,
	:email
);`

func (s *Storage) InsertDelivery(delivery *models.Delivery) error {
	tx, err := s.db.BeginTxx(context.TODO(), nil)
	if err != nil {
		return nil
	}
	defer tx.Rollback()
	if _, err := tx.NamedExec(insertDeliverySQL, delivery); err != nil {
		return err
	}
	return tx.Commit()
}

var insertItemSQL string = `
INSERT INTO item (
	order_uid,
	chrt_id,
	track_number,
	price,
	rid,
	name,
	sale,
	size,
	total_price,
	nm_id,
	brand,
	status
) VALUES (
	:order_uid,
	:chrt_id,
	:track_number,
	:price,
	:rid,
	:name,
	:sale,
	:size,
	:total_price,
	:nm_id,
	:brand,
	:status
);`

func (s *Storage) InsertItem(item *models.Item) error {
	tx, err := s.db.BeginTxx(context.TODO(), nil)
	if err != nil {
		return nil
	}
	defer tx.Rollback()
	if _, err := tx.NamedExec(insertItemSQL, item); err != nil {
		return err
	}
	return tx.Commit()
}
