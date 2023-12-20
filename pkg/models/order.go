package models

import "github.com/google/uuid"

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
