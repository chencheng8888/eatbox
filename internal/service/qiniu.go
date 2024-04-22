package service

import (
	"eat_box/global"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func (svc *Service) GetQNToken() string {
	var maxInt uint64 = 1 << 32
	putPolicy := storage.PutPolicy{
		Scope:   global.QiNiuSetting.Bucket,
		Expires: maxInt,
	}
	mac := qbox.NewMac(global.QiNiuSetting.AccessKey, global.QiNiuSetting.SecretKey)
	QNToken := putPolicy.UploadToken(mac)
	return QNToken
}
func (svc *Service) GenerateURL(key string) string {
	URL := fmt.Sprintf("http://%s/%s", global.QiNiuSetting.Domain, key)
	return URL
}
