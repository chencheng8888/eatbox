package dao

import (
	"eat_box/internal/model"
	"errors"
	"gorm.io/gorm"
)

func (d *Dao) CreateUser(user model.User) error {
	return d.engine.Create(&user).Error
}
func (d *Dao) FindUserByOpenID(openid string) (model.User, bool) {
	user := model.NewUser()
	err := d.engine.Model(&model.User{}).Where(&model.User{OpenID: openid}).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, false
	}
	return user, true
}
