package swagger

type LoginData struct {
	Token string `json:"token"`
	First bool   `json:"first"`
}
type LoginSwagger struct {
	Code string    `json:"code"`
	Msg  string    `json:"msg"`
	Data LoginData `json:"data"`
}

func NewLoginData(token string, first bool) LoginData {
	return LoginData{
		Token: token,
		First: first,
	}
}
