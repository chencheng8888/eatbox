package cache

import "github.com/go-redis/redis"

type Cache struct {
	redisdb *redis.Client
}

func NewCache(redisdb *redis.Client) *Cache {
	return &Cache{
		redisdb: redisdb,
	}
}
