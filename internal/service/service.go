package service

import (
	"eat_box/global"
	"eat_box/internal/cache"
	"eat_box/internal/dao"
)

type Service struct {
	dao   *dao.Dao
	cache *cache.Cache
}

func NewService() Service {
	return Service{
		dao:   dao.NewDao(global.DBEngine),
		cache: cache.NewCache(global.RedisDb),
	}
}
