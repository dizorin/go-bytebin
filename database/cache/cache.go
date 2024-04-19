package cache

import (
	"context"
	"github.com/dizorin/go-bytebin/models"
	"github.com/dizorin/go-bytebin/utils"
	"github.com/gofiber/fiber/v2/log"
	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

func Setup(ctx context.Context) {
	Cache = cache.New(utils.GetenvDuration("CACHE_EXPIRY"), utils.GetenvDuration("CACHE_CLEANUP"))
	Cache.OnEvicted(func(key string, _ interface{}) {
		log.Info("Cache evicted ", "key=", key)
	})
}

func Set(key string, value *utils.CompletableFuture[*models.Content]) {
	Cache.Set(key, value, cache.DefaultExpiration)
}

func Get(key string) (*utils.CompletableFuture[*models.Content], bool) {
	val, found := Cache.Get(key)
	if val == nil {
		return nil, false
	}

	return val.(*utils.CompletableFuture[*models.Content]), found
}

func Update(key string, cf *utils.CompletableFuture[*models.Content]) {
	err := Cache.Replace(key, cf, cache.DefaultExpiration)
	if err != nil {
		log.Warn(err)
	}
}
