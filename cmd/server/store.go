package main

import (
	"errors"

	pb "github.com/yawkar/wbl0/pkg/proto"
	"github.com/yawkar/wbl0/pkg/storage"
)

func deepInsertOrder(store *storage.Storage, order *pb.Order) error {
	if store == nil {
		errors.New("Store is nil!")
	}
	if order == nil {
		errors.New("Order is nil!")
	}

	orderM, err := mapOrderToModel(order)
	if err != nil {
		return err
	}
	paymentM, err := mapPaymentToModel(order)
	if err != nil {
		return err
	}
	deliveryM, err := mapDeliveryToModel(order)
	if err != nil {
		return err
	}
	itemsM, err := mapItemsToModels(order)
	if err != nil {
		return err
	}

	tx, err := store.BeginTx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = store.InsertOrder(orderM)
	if err != nil {
		return err
	}
	err = store.InsertPayment(paymentM)
	if err != nil {
		return err
	}
	err = store.InsertDelivery(deliveryM)
	if err != nil {
		return err
	}
	for _, itemM := range itemsM {
		if itemM == nil {
			return errors.New("One of items is nil!")
		}
		err := store.InsertItem(itemM)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
