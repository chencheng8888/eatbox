package config

import (
	"eat_box/global"
	"eat_box/pkg/setting"
	"errors"
	"github.com/IBM/sarama"
	"github.com/go-redis/redis"
)

func Initgorm() error {
	var err error
	global.DBEngine, err = setting.NewDBEngine(global.MYSQLsetting)
	if err != nil {
		return err
	}
	return nil
}
func Initredis() error {
	global.RedisDb = redis.NewClient(&redis.Options{
		Addr:     global.RedisSetting.Addr,
		Password: global.RedisSetting.Password,
		DB:       0,
	})
	_, err := global.RedisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func InitKafka() error {
	global.KafkaConfig = sarama.NewConfig()
	global.KafkaConfig.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	global.KafkaConfig.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	global.KafkaConfig.Consumer.Return.Errors = true
	global.KafkaConfig.Consumer.Fetch.Max = 20 // 设置一次获取消息的最大数量
	global.KafkaConfig.Producer.RequiredAcks = sarama.NoResponse
	return nil
}
func SetupSetting() error {
	set, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = set.ReadSection("mysql", &global.MYSQLsetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("redis", &global.RedisSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("wechat", &global.WechatSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("qin", &global.QiNiuSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("kafka", &global.KafkaSetting)
	if err != nil {
		return err
	}
	return nil
}
func Init() error {
	//读取配置信息
	err := SetupSetting()
	if err != nil {
		return errors.New("读取配置信息失败")
	}
	//初始化mysql数据库
	err = Initgorm()
	if err != nil {
		return errors.New("初始化mysql数据库失败")
	}
	//初始化redis数据库
	err = Initredis()
	if err != nil {
		return errors.New("初始化redis失败")
	}
	//初始化kafka
	err = InitKafka()
	if err != nil {
		return errors.New("初始化kafka失败")
	}
	return nil
}
