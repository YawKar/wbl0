package models

import (
	"github.com/google/uuid"
)

type Order struct {
	OrderUid          uuid.UUID
	TrackNumber       string
	Entry             string
	Locale            string
	InternalSignature string
	CustomerId        string
	DeliveryService   string
	ShardKey          string
	SmId              int64
	DateCreated       int64
}

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

type Delivery struct {
	Id       int64
	OrderUid uuid.UUID
	Name     string
	Phone    string
	Zip      string
	City     string
	Address  string
	Region   string
	Email    string
}

type Item struct {
	Id          int64
	OrderUid    uuid.UUID
	ChrtId      int64
	TrackNumber string
	Price       int64
	Rid         uuid.UUID
	Name        string
	Sale        int64
	Size        string
	TotalPrice  int64
	NmId        int64
	Brand       string
	Status      int32
}
