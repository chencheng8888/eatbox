package model

type User struct {
	ID        string `gorm:"column:id;primaryKey" json:"id"`
	NickName  string `gorm:"column:nick_name" json:"nick_name"`
	HeadImage string `gorm:"column:head_image" json:"head_image"`
	Points    int    `gorm:"column:points" json:"points"`
	Level     int    `gorm:"column:level;default:0" json:"level"`
	Tele      string `gorm:"column:tele" json:"tele"`
	OpenID    string `gorm:"column:open_id" json:"open_id"`
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
