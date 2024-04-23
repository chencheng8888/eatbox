package swagger

type LoginData struct {
	Token string `json:"token"` //token
	First bool   `json:"first"` //是否第一次登录，如果是就是true,否则就false
}
type LoginSwagger struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data LoginData `json:"data"`
}
type UpdateSwagger struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type DetailData struct {
	ID        string `json:"id"`         //id
	NickName  string `json:"nick_name"`  //昵称
	HeadImage string `json:"head_image"` //头像
	Points    int    `json:"points"`     //积分
	Level     int    `json:"level"`      //等级
	Tele      string `json:"tele"`       //电话
}
type DetailSwagger struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data DetailData `json:"data"`
}

func NewLoginData(token string, first bool) LoginData {
	return LoginData{
		Token: token,
		First: first,
	}
}
