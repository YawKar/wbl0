package main

import (
	"errors"

	"github.com/google/uuid"
	"github.com/yawkar/wbl0/pkg/models"
	pb "github.com/yawkar/wbl0/pkg/proto"
)

func mapOrderToModel(order *pb.Order) (*models.Order, error) {
	if order == nil {
		return nil, errors.New("Order is nil!")
	}
	model := &models.Order{}
	if orderUid, err := uuid.Parse(order.OrderUid); err != nil {
		return nil, err
	} else {
		model.OrderUid = orderUid
	}

	model.TrackNumber = order.TrackNumber
	model.Entry = order.Entry
	model.Locale = order.Locale
	model.InternalSignature = order.InternalSignature
	model.CustomerId = order.CustomerId
	model.DeliveryService = order.DeliveryService
	model.ShardKey = order.ShardKey
	model.SmId = order.SmId
	model.DateCreated = order.DateCreated.AsTime().Unix()

	return model, nil
}

func mapPaymentToModel(order *pb.Order) (*models.Payment, error) {
	// TODO
}

func mapDeliveryToModel(order *pb.Order) (*models.Delivery, error) {
	// TODO
}

func mapItemsToModels(order *pb.Order) ([]*models.Item, error) {
	// TODO
}

func mapItemToModel(order *pb.Order, item *pb.Item) (*models.Item, error) {
	// TODO
}
