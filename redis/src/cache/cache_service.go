package cache

import "github.com/redis/go-redis/v9"

var RC *redis.Client

type Cache interface {
	StartCacheClient()
}

type Connector struct {
	CacheService Cache
}

func (c *Connector) StartCacheClientConnection() {
	c.CacheService.StartCacheClient()
}
