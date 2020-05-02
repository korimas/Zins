package cache

import (
	"github.com/goburrow/cache"
	"github.com/zpdev/zins/model"
	"github.com/zpdev/zins/product/extend"
	"time"
)

var TokenCache = NewTokenCache()

type tokenCache struct {
	cache cache.LoadingCache
}

func (c *tokenCache) Get(tokenID string) *model.Token {
	val, err := c.cache.Get(tokenID)
	if err != nil {
		return nil
	}
	if val != nil {
		return val.(*model.Token)
	}
	return nil
}

func NewTokenCache() *tokenCache {
	load := func(k cache.Key) (value cache.Value, err error) {
		token := &model.Token{}
		if extend.DB().Where("Token = ?", k.(string)).First(token).RecordNotFound() {
			value = nil
		} else {
			value = token
		}
		return
	}
	newCache := cache.NewLoadingCache(
		load,
		cache.WithMaximumSize(100),
		cache.WithExpireAfterAccess(30*time.Minute),
	)
	return &tokenCache{
		cache: newCache,
	}
}
