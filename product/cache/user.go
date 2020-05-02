package cache

import (
	"github.com/goburrow/cache"
	"github.com/zpdev/zins/model"
	"github.com/zpdev/zins/product/extend"
	"time"
)

var UserCache = NewUserCache()

type userCache struct {
	cache cache.LoadingCache
}

func (c *userCache) Get(username string) *model.User {
	val, err := c.cache.Get(username)
	if err != nil {
		return nil
	}
	if val != nil {
		return val.(*model.User)
	}
	return nil
}

func NewUserCache() *userCache {
	load := func(k cache.Key) (value cache.Value, err error) {
		user := &model.User{}
		if extend.DB().Where("Username = ?", k.(string)).First(user).RecordNotFound() {
			value = nil
		} else {
			value = user
		}
		return
	}
	newCache := cache.NewLoadingCache(
		load,
		cache.WithMaximumSize(100),
		cache.WithExpireAfterAccess(30*time.Minute),
	)
	return &userCache{
		cache: newCache,
	}
}
