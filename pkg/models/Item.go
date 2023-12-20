package models

import (
	"github.com/google/uuid"
)

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
