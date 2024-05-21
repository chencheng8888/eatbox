package service

import (
	"eat_box/global"
	"eat_box/internal/cache"
	"eat_box/internal/dao"
	"eat_box/internal/kafka"
)

type Service struct {
	dao      *dao.Dao
	cache    *cache.Cache
	producer kafka.Producer
}

func NewService() Service {
	return Service{
		dao:      dao.NewDao(global.DBEngine),
		cache:    cache.NewCache(global.RedisDb),
		producer: kafka.NewProducer([]string{global.KafkaSetting.Addr}, global.KafkaConfig),
	}
}
