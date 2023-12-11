package cache

import (
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"tech-stack/redis/config"
)

var RedisClient *redis.Client

func StartCacheClient() {
	db, err := strconv.Atoi(config.Get("REDIS_DB"))
	if err != nil {
		log.Fatal("Cache resource not found.")
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Get("REDIS_ADDR"),
		Password: config.Get("REDIS_PASS"),
		DB:       db,
	})
}