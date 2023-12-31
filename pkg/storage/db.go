package storage

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/yawkar/wbl0/pkg/models"
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

func insertOrder(db *sqlx.DB, order *models.Order) error {
	tx, err := db.BeginTxx(context.TODO(), nil)
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

func insertPayment(db *sqlx.DB, payment *models.Payment) error {
	tx, err := db.BeginTxx(context.TODO(), nil)
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

func insertDelivery(db *sqlx.DB, delivery *models.Delivery) error {
	tx, err := db.BeginTxx(context.TODO(), nil)
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

func insertItem(db *sqlx.DB, item *models.Item) error {
	tx, err := db.BeginTxx(context.TODO(), nil)
	if err != nil {
		return nil
	}
	defer tx.Rollback()
	if _, err := tx.NamedExec(insertItemSQL, item); err != nil {
		return err
	}
	return tx.Commit()
}

var selectOrderSQL string = `
SELECT *
  FROM "order"
 WHERE order_uid = $1;`

func getOrder(db *sqlx.DB, orderUid uuid.UUID) (*models.Order, error) {
	tx, err := db.BeginTxx(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	m := &models.Order{}
	if err := tx.Get(m, selectOrderSQL, orderUid); err != nil {
		return nil, err
	}
	return m, nil
}

var selectPaymentSQL string = `
SELECT *
  FROM payment
 WHERE transaction = $1;`

func getPayment(db *sqlx.DB, orderUid uuid.UUID) (*models.Payment, error) {
	tx, err := db.BeginTxx(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	m := &models.Payment{}
	if err := tx.Get(m, selectPaymentSQL, orderUid); err != nil {
		return nil, err
	}
	return m, nil
}

var selectDeliverySQL string = `
SELECT *
  FROM delivery
 WHERE order_uid = $1;`

func getDelivery(db *sqlx.DB, orderUid uuid.UUID) (*models.Delivery, error) {
	tx, err := db.BeginTxx(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	m := &models.Delivery{}
	if err := tx.Get(m, selectDeliverySQL, orderUid); err != nil {
		return nil, err
	}
	return m, nil
}

var selectItemsSQL string = `
SELECT *
  FROM item
 WHERE order_uid = $1;`

func getItems(db *sqlx.DB, orderUid uuid.UUID) ([]*models.Item, error) {
	tx, err := db.BeginTxx(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	m := make([]*models.Item, 0)
	if err := tx.Select(&m, selectItemsSQL, orderUid); err != nil {
		return nil, err
	}
	return m, nil
}

var selectFirstNOrdersSQL string = `
SELECT *
  FROM "order"
 LIMIT $1;`

func getRandomOrders(db *sqlx.DB, amount int) ([]*models.Order, error) {
	tx, err := db.BeginTxx(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	m := make([]*models.Order, 0)
	if err := tx.Select(&m, selectFirstNOrdersSQL, amount); err != nil {
		return nil, err
	}
	return m, nil
}
