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
	AppID     string `json:"Appid"`
	AppSecret string `json:"Appsecret"`
}
type QiniuSettingS struct {
	AccessKey string `json:"AccessKey"`
	SecretKey string `json:"SecretKey"`
	Bucket    string `json:"Bucket"`
	Domain    string `json:"Domain"`
}
type KafkaSettings struct {
	Addr string `json:"addr"`
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
