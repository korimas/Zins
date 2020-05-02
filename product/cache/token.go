package cache

import (
	"github.com/goburrow/cache"
	"github.com/zpdev/zins/model"
)

type TokenCache struct {
	cache cache.LoadingCache
}

func (token *TokenCache) Get(tokenID string) *model.Token {

}
