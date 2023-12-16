package main

import (
	"log"
	"tech-stack/redis/src/cache"
	"tech-stack/redis/src/service"
)

func main() {

	cacheService := cache.Connector{CacheService: &cache.RedisClient{}}
	cacheService.StartCacheClientConnection()

	write := service.Write("domain", "daddydemir.dev")
	if write {
		log.Println("writing is successfully.")
	}

	read := service.Read("domain")
	log.Println("domain:", read)
}
