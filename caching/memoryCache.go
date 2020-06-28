package caching

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type memoryCache struct {
	Cache *cache.Cache
}

type IMemoryCache interface {
	Get(key string) (interface{}, bool)
	Set(key string, data interface{})
}

func New() IMemoryCache {
	return &memoryCache{Cache: cache.New(8*time.Hour, 10*time.Hour)}
}

func (c *memoryCache) Get(key string) (interface{}, bool) {
	data, isExist := c.Cache.Get(key)
	if !isExist {
		return nil, false
	}

	return data, true
}

func (c *memoryCache) Set(key string, data interface{}) {
	c.Cache.Set(key, data, time.Hour*12)
}
