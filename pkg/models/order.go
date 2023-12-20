package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderUid          uuid.UUID `db:"order_uid"`
	TrackNumber       string    `db:"track_number"`
	Entry             string    `db:"entry"`
	Locale            string    `db:"locale"`
	InternalSignature string    `db:"internal_signature"`
	CustomerId        string    `db:"customer_id"`
	DeliveryService   string    `db:"delivery_service"`
	ShardKey          string    `db:"shardkey"`
	SmId              int64     `db:"sm_id"`
	DateCreated       time.Time `db:"date_created"`
	OofShard          string    `db:"oof_shard"`
}
