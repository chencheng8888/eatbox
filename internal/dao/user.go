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
		ok := d.IsExistUser(id) //检验id是否存在
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
func (d *Dao) UpdateUserInfo(id string, mp map[string]interface{}) error {
	return d.engine.Model(&model.User{}).Where("id = ?", id).Updates(mp).Error
}

func (d *Dao) GetBusinessNum() int64 {
	var num int64
	d.engine.Model(&model.Business{}).Count(&num)
	return num
}
func (d *Dao) GetBusinesses(limit int, offset int) ([]model.Business, error) {
	var busses []model.Business
	err := d.engine.Model(&model.Business{}).Limit(limit).Offset(offset).Find(&busses).Error
	return busses, err
}
func (d *Dao) IsExistUser(id string) bool {
	var num int64
	d.engine.Model(&model.User{}).Where("id = ?",id).Count(&num)
	return num > 0
}