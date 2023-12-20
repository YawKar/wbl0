package main

import (
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"github.com/yawkar/wbl0/pkg/models"
	pb "github.com/yawkar/wbl0/pkg/proto"
)

func mapOrderToModel(order *pb.Order) (*models.Order, error) {
	if order == nil {
		return nil, errors.New("order is nil")
	}
	orderUid, err := uuid.Parse(order.OrderUid)
	if err != nil {
		return nil, err
	}
	return &models.Order{
		OrderUid:          orderUid,
		TrackNumber:       order.TrackNumber,
		Entry:             order.Entry,
		Locale:            order.Locale,
		InternalSignature: order.InternalSignature,
		CustomerId:        order.CustomerId,
		DeliveryService:   order.DeliveryService,
		ShardKey:          order.ShardKey,
		SmId:              order.SmId,
		DateCreated:       order.DateCreated.AsTime(),
		OofShard:          order.OofShard,
	}, nil
}

func mapPaymentToModel(order *pb.Order) (*models.Payment, error) {
	if order == nil {
		return nil, errors.New("order is nil")
	}
	if order.Payment == nil {
		return nil, errors.New("payment is nil")
	}
	orderUuid, err := uuid.Parse(order.OrderUid)
	if err != nil {
		return nil, err
	}
	var requestId uuid.NullUUID
	if reqUuid, err := uuid.Parse(order.Payment.RequestId); err != nil {
		requestId.UUID = reqUuid
		requestId.Valid = true
	}
	return &models.Payment{
		Id:           0,
		Transaction:  orderUuid,
		RequestId:    requestId,
		Currency:     order.Payment.Currency,
		Provider:     order.Payment.Provider,
		Amount:       order.Payment.Amount,
		PaymentDt:    order.Payment.PaymentDt,
		Bank:         order.Payment.Bank,
		DeliveryCost: order.Payment.DeliveryCost,
		GoodsTotal:   order.Payment.GoodsTotal,
		CustomFee:    order.Payment.CustomFee,
	}, nil
}

func mapDeliveryToModel(order *pb.Order) (*models.Delivery, error) {
	if order == nil {
		return nil, errors.New("order is nil")
	}
	if order.Delivery == nil {
		return nil, errors.New("delivery is nil")
	}
	orderUuid, err := uuid.Parse(order.OrderUid)
	if err != nil {
		return nil, err
	}
	return &models.Delivery{
		Id:       0,
		OrderUid: orderUuid,
		Name:     order.Delivery.Name,
		Phone:    order.Delivery.Phone,
		Zip:      order.Delivery.Zip,
		City:     order.Delivery.City,
		Address:  order.Delivery.Address,
		Region:   order.Delivery.Region,
		Email:    order.Delivery.Email,
	}, nil
}

func mapItemsToModels(order *pb.Order) ([]*models.Item, error) {
	if order == nil {
		return nil, errors.New("order is nil")
	}
	if order.Items == nil {
		slog.Debug("order.Items is nil")
	}

	itemsM := make([]*models.Item, len(order.Items))
	orderUuid, err := uuid.Parse(order.OrderUid)
	if err != nil {
		return nil, err
	}
	for i, itemP := range order.Items {
		if itemsM[i], err = mapItemToModel(orderUuid, itemP); err != nil {
			return nil, errors.Join(errors.New("one of order items is broken"), err)
		}
	}
	return itemsM, nil
}

func mapItemToModel(orderUuid uuid.UUID, itemP *pb.Item) (*models.Item, error) {
	if itemP == nil {
		return nil, errors.New("item is nil")
	}
	rid, err := uuid.Parse(itemP.Rid)
	if err != nil {
		return nil, errors.Join(errors.New("item has incorrect rid"), err)
	}
	return &models.Item{
		Id:          0,
		OrderUid:    orderUuid,
		ChrtId:      itemP.ChrtId,
		TrackNumber: itemP.TrackNumber,
		Price:       itemP.Price,
		Rid:         rid,
		Name:        itemP.Name,
		Sale:        itemP.Sale,
		Size:        itemP.Size,
		TotalPrice:  itemP.TotalPrice,
		NmId:        itemP.NmId,
		Brand:       itemP.Brand,
		Status:      itemP.Status,
	}, nil
}
