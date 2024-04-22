package dao

import (
	"eat_box/internal/model"
	"eat_box/pkg/app/random"
	"errors"
	"gorm.io/gorm"
)

func (d *Dao) CreateUser(user model.User) error {
	var id string
	for {
		id = random.GenerateRandomNumberString(10)
		_, ok := d.FindUserByID(id)
		if !ok {
			break
		}
	}
	user.ID = id
	return d.engine.Create(&user).Error
}
func (d *Dao) FindUserByID(id string) (model.User, bool) {
	user := model.NewUser()
	err := d.engine.Model(&model.User{}).Where(&model.User{ID: id}).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, false
	}
	return user, true
}
func (d *Dao) FindUserByOpenID(openid string) (model.User, bool) {
	user := model.NewUser()
	err := d.engine.Model(&model.User{}).Where(&model.User{OpenID: openid}).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, false
	}
	return user, true
}
func (d *Dao) UpdateUserInfo(user model.User) error {
	return d.engine.Updates(&user).Error
}
