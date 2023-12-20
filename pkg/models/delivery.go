package models

import "github.com/google/uuid"

type Delivery struct {
	Id       int64			`db:"id"`
	OrderUid uuid.UUID	`db:"order_uid"`
	Name     string			`db:"name"`
	Phone    string			`db:"phone"`
	Zip      string			`db:"zip"`
	City     string			`db:"city"`
	Address  string			`db:"address"`
	Region   string			`db:"region"`
	Email    string			`db:"email"`
}
