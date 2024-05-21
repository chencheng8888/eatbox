package global

import (
	"eat_box/pkg/setting"
	"time"
)

var (
	MYSQLsetting  *setting.MysqlSettings
	RedisSetting  *setting.RedisSettingS
	WechatSetting *setting.WechatSettings
	QiNiuSetting  *setting.QiniuSettingS
	KafkaSetting  *setting.KafkaSettings
	//JWTSetting   *setting.Jwtsettings
)

// 定义过期时间
const TokenExpireDuration = time.Hour * 24

var Secret = []byte("eatbox")
