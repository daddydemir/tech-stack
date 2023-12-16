package service

import (
	"context"
	"log"
	"tech-stack/redis/src/cache"
	"time"
)

func Write(key string, data interface{}) bool {
	status := cache.RC.Ping(context.Background()).Err()
	if status != nil {
		log.Println("Cache connection is unreachable.")
		return false
	}

	response := cache.RC.Set(context.Background(), key, data, 10*time.Hour)
	if response.Err() != nil {
		log.Println("data writing for cache server has occurred error:", response.Err().Error())
		return false
	}

	return true
}

func Read(key string) interface{} {
	status := cache.RC.Ping(context.Background()).Err()
	if status != nil {
		log.Println("Cache connection is unreachable.")
		return false
	}

	result := cache.RC.Get(context.Background(), key)
	if result.Err() != nil {
		log.Println("data reading for cache server has occurred error:", result.Err().Error())
	}

	return result.Val()
}
