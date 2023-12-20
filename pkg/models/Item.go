package models

import (
	"github.com/google/uuid"
)

type Item struct {
	Id          int64			`db:"id"`
	OrderUid    uuid.UUID	`db:"order_uid"`
	ChrtId      int64			`db:"chrt_id"`
	TrackNumber string		`db:"track_number"`
	Price       int64			`db:"price"`
	Rid         uuid.UUID	`db:"rid"`
	Name        string		`db:"name"`
	Sale        int64			`db:"sale"`
	Size        string		`db:"size"`
	TotalPrice  int64			`db:"total_price"`
	NmId        int64			`db:"nm_id"`
	Brand       string		`db:"brand"`
	Status      int32			`db:"status"`
}
