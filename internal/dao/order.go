package dao

import (
	"eat_box/internal/model"
)

func (d *Dao) CreateOrder(order model.Order) error {
	tx := d.engine.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err := tx.Create(&order).Error
	if err != nil {
		tx.Rollback()
	}
	oau := model.NewOAU(order.UserID, order.Model.ID)
	err = tx.Create(&oau).Error
	if err != nil {
		tx.Rollback()
	}
	oab := model.NewOAB(order.BusID, order.Model.ID)
	err = tx.Create(&oab).Error
	if err != nil {
		tx.Rollback()
	}
	return tx.Commit().Error
}

func (d *Dao) UpdateOrder(id uint, mp map[string]interface{}) error {
	return d.engine.Model(model.Order{}).Where("id = ?", id).Updates(mp).Error
}
func (d *Dao) FindOrderByID(id uint) (model.Order, error) {
	var order model.Order
	err := d.engine.Find(&order, id).Error
	return order, err
}
func (d *Dao) CreateOAU(oau model.OrderAndUserID) error {
	return d.engine.Create(&oau).Error
}
func (d *Dao) FindOAUByUserID(uid string) ([]uint, error) {
	var orders []uint
	err := d.engine.Model(&model.OrderAndUserID{}).Where(&model.OrderAndUserID{UserID: uid}).Find(&orders).Error
	return orders, err
}
func (d *Dao) CreateOAB(oab model.OrderAndBusID) error {
	return d.engine.Create(&oab).Error
}
func (d *Dao) FindOAUByBusID(bid uint) ([]uint, error) {
	var orders []uint
	err := d.engine.Model(&model.OrderAndBusID{}).Where(&model.OrderAndBusID{BusID: bid}).Find(&orders).Error
	return orders, err
}
