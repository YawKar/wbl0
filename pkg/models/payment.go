package models

import "github.com/google/uuid"

type Payment struct {
	Id           int64					`db:"id"`
	Transaction  uuid.UUID			`db:"transaction"`
	RequestId    uuid.NullUUID	`db:"request_id"`
	Currency     string					`db:"currency"`
	Provider     string					`db:"provider"`
	Amount       int64					`db:"amount"`
	PaymentDt    int64					`db:"payment_dt"`
	Bank         string					`db:"bank"`
	DeliveryCost int64					`db:"delivery_cost"`
	GoodsTotal   int64					`db:"goods_total"`
	CustomFee    int64					`db:"custom_fee"`
}
