package systems

import (
	"github.com/allegro/bigcache"
	"sync"
	"time"
)

type CacheSystem struct {
	bigCache *bigcache.BigCache
}

var instance *CacheSystem
var once sync.Once

func GetCacheSystemInstance() *CacheSystem {
	once.Do(func() {
		cacheConfig := bigcache.DefaultConfig(5 * time.Minute)
		cacheConfig.CleanWindow = 6 * time.Minute
		cache, _ := bigcache.NewBigCache(cacheConfig)
		instance = &CacheSystem{
			bigCache: cache,
		}
	})
	return instance
}

func (cs *CacheSystem) Set(key string, value []byte) error {
	err := cs.bigCache.Set(key, value)

	if err != nil {
		return err
	}
	return nil
}

func (cs *CacheSystem) Get(key string) ([]byte, error) {
	return cs.bigCache.Get(key)
}

func (cs *CacheSystem) Delete(key string) error {
	return cs.bigCache.Delete(key)
}
