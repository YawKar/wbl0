package storage

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"github.com/yawkar/wbl0/pkg/models"
)

type cacheConfig struct {
	CacheExpiration time.Duration
	CleanupInterval time.Duration
}

func mkCache(c *cacheConfig) (*cache.Cache, error) {
	cache := cache.New(c.CacheExpiration, c.CleanupInterval)
	return cache, nil
}

func toCacheOrderKey(order *models.Order) string {
	return fmt.Sprintf("%s:order", order.OrderUid)
}

func toCachePaymentKey(payment *models.Payment) string {
	return fmt.Sprintf("%s:payment", payment.Transaction)
}

func toCacheDeliveryKey(delivery *models.Delivery) string {
	return fmt.Sprintf("%s:delivery", delivery.OrderUid)
}

func toCacheItemKey(item *models.Item) string {
	return fmt.Sprintf("%s:item:%d", item.OrderUid, item.ChrtId)
}

func toCacheOrderItemsListKey(orderUuid uuid.UUID) string {
	return fmt.Sprintf("%s:itemschrts", orderUuid)
}

func cacheOrder(cache *cache.Cache, order *models.Order) {
	cache.SetDefault(toCacheOrderKey(order), order)
}

func cachePayment(cache *cache.Cache, payment *models.Payment) {
	cache.SetDefault(toCachePaymentKey(payment), payment)
}

func cacheDelivery(cache *cache.Cache, delivery *models.Delivery) {
	cache.SetDefault(toCacheDeliveryKey(delivery), delivery)
}

func cacheItem(cache *cache.Cache, item *models.Item) {
	itemsListKey := toCacheOrderItemsListKey(item.OrderUid)
	itemsList := make([]string, 0)
	if items, ok := cache.Get(itemsListKey); !ok {
	} else if casted, ok := items.([]string); !ok {
		slog.Debug("couldn't cast cached items to []string. resetting cached itemsListKey", "orderUid", item.OrderUid)
		cache.Delete(itemsListKey)
	} else {
		itemsList = casted
	}
	itemsList = append(itemsList, fmt.Sprintf("%d", item.ChrtId))
	cache.SetDefault(itemsListKey, itemsList)
	cache.SetDefault(toCacheItemKey(item), item)
}
