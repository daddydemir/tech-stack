package cache

import (
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"tech-stack/redis/config"
)

type RedisClient struct {
}

func (r *RedisClient) StartCacheClient() {
	db, err := strconv.Atoi(config.Get("REDIS_DB"))
	if err != nil {
		log.Fatal("Cache resource not found.")
	}

	RC = redis.NewClient(&redis.Options{
		Addr:     config.Get("REDIS_ADDR"),
		Password: config.Get("REDIS_PASS"),
		DB:       db,
	})
}
