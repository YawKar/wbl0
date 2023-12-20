package models

import "github.com/google/uuid"

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
