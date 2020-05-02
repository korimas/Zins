package cache

import (
	"github.com/goburrow/cache"
	"github.com/zpdev/zins/model"
)

type UserCache struct {
	cache cache.LoadingCache
}

func (user *UserCache) Get(username string) *model.User {

}
