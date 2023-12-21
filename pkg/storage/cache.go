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

func toCacheOrderKey(orderUid uuid.UUID) string {
	return fmt.Sprintf("%s:order", orderUid)
}

func toCachePaymentKey(orderUid uuid.UUID) string {
	return fmt.Sprintf("%s:payment", orderUid)
}

func toCacheDeliveryKey(orderUid uuid.UUID) string {
	return fmt.Sprintf("%s:delivery", orderUid)
}

func toCacheItemKey(orderUid uuid.UUID, chrtId int64) string {
	return fmt.Sprintf("%s:item:%d", orderUid, chrtId)
}

func toCacheOrderItemsListKey(orderUuid uuid.UUID) string {
	return fmt.Sprintf("%s:itemschrts", orderUuid)
}

func cacheOrder(cache *cache.Cache, order *models.Order) {
	cache.SetDefault(toCacheOrderKey(order.OrderUid), order)
}

func cachePayment(cache *cache.Cache, payment *models.Payment) {
	cache.SetDefault(toCachePaymentKey(payment.Transaction), payment)
}

func cacheDelivery(cache *cache.Cache, delivery *models.Delivery) {
	cache.SetDefault(toCacheDeliveryKey(delivery.OrderUid), delivery)
}

func cacheItem(cache *cache.Cache, item *models.Item) {
	itemsListKey := toCacheOrderItemsListKey(item.OrderUid)
	itemsList := make([]int64, 0)
	if items, ok := cache.Get(itemsListKey); !ok {
	} else if casted, ok := items.([]int64); !ok {
		slog.Debug("couldn't cast cached items to []int64. resetting cached itemsListKey", "cacheKey", itemsListKey)
		cache.Delete(itemsListKey)
	} else {
		itemsList = casted
	}
	itemsList = append(itemsList, item.ChrtId)
	cache.SetDefault(itemsListKey, itemsList)
	cache.SetDefault(toCacheItemKey(item.OrderUid, item.ChrtId), item)
}

func getCachedOrder(cache *cache.Cache, orderUid uuid.UUID) (m *models.Order, found bool) {
	cacheKey := toCacheOrderKey(orderUid)
	if cached, found := cache.Get(cacheKey); !found {
		return nil, false
	} else if m, ok := cached.(*models.Order); !ok {
		slog.Debug("couldn't cast cached order to *model. resetting cache key", "cacheKey", cacheKey)
		cache.Delete(cacheKey)
		return nil, false
	} else {
		return m, true
	}
}

func getCachedPayment(cache *cache.Cache, orderUid uuid.UUID) (m *models.Payment, found bool) {
	cacheKey := toCachePaymentKey(orderUid)
	if cached, found := cache.Get(cacheKey); !found {
		return nil, false
	} else if m, ok := cached.(*models.Payment); !ok {
		slog.Debug("couldn't cast cached payment to *model. resetting cache key", "cacheKey", cacheKey)
		cache.Delete(cacheKey)
		return nil, false
	} else {
		return m, true
	}
}

func getCachedDelivery(cache *cache.Cache, orderUid uuid.UUID) (m *models.Delivery, found bool) {
	cacheKey := toCacheDeliveryKey(orderUid)
	if cached, found := cache.Get(cacheKey); !found {
		return nil, false
	} else if m, ok := cached.(*models.Delivery); !ok {
		slog.Debug("couldn't cast cached delivery to *model. resetting cache key", "cacheKey", cacheKey)
		cache.Delete(cacheKey)
		return nil, false
	} else {
		return m, true
	}
}

func getCachedItems(cache *cache.Cache, orderUid uuid.UUID) (m []*models.Item, found bool) {
	chrtIdsCacheKey := toCacheOrderItemsListKey(orderUid)
	if cached, found := cache.Get(chrtIdsCacheKey); !found {
		return nil, false
	} else if chrtIds, ok := cached.([]int64); !ok {
		slog.Debug("couldn't cast cached chrtIds to []int64. resetting cache key", "cacheKey", chrtIdsCacheKey)
		cache.Delete(chrtIdsCacheKey)
		return nil, false
	} else {
		m = make([]*models.Item, 0, len(chrtIds))
		for _, chrtId := range chrtIds {
			if itemM, ok := getCachedItem(cache, orderUid, chrtId); !ok {
				slog.Debug("cache miss: one of items wasn't found", "orderUid", orderUid, "chrtId", chrtId)
				return nil, false
			} else {
				m = append(m, itemM)
			}
		}
		return m, true
	}
}

func getCachedItem(cache *cache.Cache, orderUid uuid.UUID, chrtId int64) (m *models.Item, found bool) {
	cacheKey := toCacheItemKey(orderUid, chrtId)
	if cached, found := cache.Get(cacheKey); !found {
		return nil, false
	} else if m, ok := cached.(*models.Item); !ok {
		slog.Debug("couldn't cast cached item to *model. resetting cache key", "cacheKey", cacheKey)
		cache.Delete(cacheKey)
		return nil, false
	} else {
		return m, true
	}
}
