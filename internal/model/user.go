package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string    `gorm:"column:id;primaryKey" json:"id"`
	NickName  string    `gorm:"column:nick_name" json:"nick_name"`
	HeadImage string    `gorm:"column:head_image" json:"head_image"`
	Points    int       `gorm:"column:points" json:"points"`
	Level     int       `gorm:"column:level;default:0" json:"level"`
	Tele      string    `gorm:"column:tele" json:"tele"`
	OpenID    string    `gorm:"column:open_id" json:"open_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdateAt  time.Time `gorm:"column:update_at" json:"update_at"`
}

func NewUser() User {
	return User{}
}
func (u *User) Create(openid string) {
	u.OpenID = openid
}
func (u *User) BeforeCreate(*gorm.DB) error {
	u.CreatedAt = time.Now()
	u.UpdateAt = time.Now()
	return nil
}
func (u *User) BeforeUpdate(*gorm.DB) error {
	u.UpdateAt = time.Now()
	return nil
}
