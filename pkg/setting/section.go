package setting

type MysqlSettings struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Dbname   string `json:"dbname"`
}
type RedisSettingS struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
}
type WechatSettings struct {
	AppID     string `json:"appid"`
	AppSecret string `json:"appsecret"`
}
type QiniuSettingS struct {
	AccessKey string `json:"AccessKey"`
	SecretKey string `json:"SecretKey"`
	Bucket    string `json:"Bucket"`
	Domain    string `json:"Domain"`
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}