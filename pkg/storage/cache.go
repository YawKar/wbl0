package storage

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type cacheConfig struct {
	CacheExpiration time.Duration
	CleanupInterval time.Duration
}

func mkCache(c *cacheConfig) (*cache.Cache, error) {
	cache := cache.New(c.CacheExpiration, c.CleanupInterval)
	return cache, nil
}
