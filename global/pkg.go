package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	DBEngine *gorm.DB
	RedisDb  *redis.Client
)
