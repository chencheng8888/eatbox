package model

type User struct {
	ID        string `gorm:"primaryKey" json:"id"`
	NickName  string `json:"nick_name"`
	HeadImage string `json:"head_image"`
	Points    int    `json:"points"`
	Level     int    `json:"level"`
	Tele      string `json:"tele"`
	OpenID    string `json:"open_id"`
}

func NewUser() User {
	return User{}
}
func (u *User) Create(openid string) {
	u.OpenID = openid
}
func (u *User) UpdateNickName(nickname string) {
	u.NickName = nickname
}
func (u *User) UpdateHeadImage(headimage string) {
	u.HeadImage = headimage
}
func (u *User) UpdatePoints(points int) {
	u.Points = points
}
func (u *User) UpdateLevel(level int) {
	u.Level = level
}
func (u *User) UpdateTele(tele string) {
	u.Tele = tele
}
