package swagger

type QiniuData struct {
	Token string `json:"token"`
}
type QiniuSwagger struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data QiniuData `json:"data"`
}
