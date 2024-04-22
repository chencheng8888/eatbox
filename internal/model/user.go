package model

type User struct {
	ID         string `gorm:"primaryKey" json:"id"`
	NickName   string `json:"nick_name"`
	HeadImage  string `json:"head_image"`
	Points     int    `json:"points"`
	Level      int    `json:"level"`
	Tele       string `json:"tele"`
	WechatName string `json:"wechat_name"`
	OpenID     string `json:"open_id"`
}

func NewUser() User {
	return User{}
}
func (u *User) Create(openid string) {
	u.OpenID = openid
}
