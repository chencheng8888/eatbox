package dao

import "eat_box/internal/model"

func (d *Dao) CreateBox(box model.BlindBox) error {
	return d.engine.Create(&box).Error
}
func (d *Dao) UpdateBox(box model.BlindBox) error {
	return d.engine.Updates(&box).Error
}
func (d *Dao) DeleteBox(box model.BlindBox) error {
	return d.engine.Delete(&box).Error
}
func (d *Dao) GetBox(id int) (model.BlindBox, error) {
	box := model.NewBlindBox()
	err := d.engine.First(&box, id).Error
	return box, err
}
func (d *Dao) CreateBusiness(Business model.Business) error {
	return d.engine.Create(&Business).Error
}
func (d *Dao) UpdateBusiness(Business model.Business) error {
	return d.engine.Updates(&Business).Error
}
func (d *Dao) DeleteBusiness(Business model.Business) error {
	return d.engine.Delete(&Business).Error
}
func (d *Dao) GetBusiness(id int) (model.Business, error) {
	Business := model.NewBusiness()
	err := d.engine.First(&Business, id).Error
	return Business, err
}
func (d *Dao) CreateBusinessAndBox(Business model.Business) error {
	return d.engine.Create(&Business).Error
}
func (d *Dao) DeleteBusinessAndBox(Business model.Business) error {
	return d.engine.Delete(&Business).Error
}
func (d *Dao) FindBusinessAndBoxByID(id int) (model.BusinessAndBox, error) {
	BusinessAndBox := model.NewBusinessAndBox()
	err := d.engine.First(&BusinessAndBox, id).Error
	return BusinessAndBox, err
}
func (d *Dao) FindBusinessAndBoxByBusID(id int) ([]model.BusinessAndBox, error) {
	var b []model.BusinessAndBox
	err := d.engine.Where(&model.BusinessAndBox{BusinessID: id}).Find(&b).Error
	return b, err
}
func (d *Dao) FindBusinessAndBoxByBoxID(id int) ([]model.BusinessAndBox, error) {
	var b []model.BusinessAndBox
	err := d.engine.Where(&model.BusinessAndBox{BoxID: id}).Find(&b).Error
	return b, err
}
