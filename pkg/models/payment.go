package models

import "github.com/google/uuid"

type Payment struct {
	Id           int64
	Transaction  uuid.UUID
	RequestId    uuid.NullUUID
	Currency     string
	Provider     string
	Amount       int64
	PaymentDt    int64
	Bank         string
	DeliveryCost int64
	GoodsTotal   int64
	CustomFee    int64
}
