package dao

import (
	"eat_box/internal/model"
	"eat_box/internal/model/swagger"
)

// CreateBusiness 创建商家
func (d *Dao) CreateBusiness(b model.Business) error {
	tx := d.engine.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err := tx.Create(&b).Error
	if err != nil {
		tx.Rollback()
	}
	score := model.BusinessScore{
		ID:        int(b.ID),
		ScoreNum:  0,
		PeopleNum: 0,
	}
	err = tx.Create(&score).Error
	if err != nil {
		tx.Rollback()
	}
	return tx.Commit().Error
}

// UpdateBusinessesScore 更新商家的评分
func (d *Dao) UpdateBusinessesScore(datas []swagger.ScoreData) error {
	tx := d.engine.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	tx.SavePoint("test1")
	for _, data := range datas {
		var business model.Business
		var businessScore model.BusinessScore
		mp1 := make(map[string]interface{})
		mp2 := make(map[string]interface{})
		err := tx.First(&business, data.BusinessID).Error
		if err != nil {
			tx.RollbackTo("test1")
		}
		err = tx.First(&businessScore, data.BusinessID).Error
		if err != nil {
			tx.RollbackTo("test1")
		}
		//businessScore.ScoreNum += data.Score
		mp1["score_num"] = businessScore.ScoreNum + data.Score
		//businessScore.PeopleNum++
		mp1["people_num"] = businessScore.PeopleNum + 1
		//business.Score = businessScore.ScoreNum / float64(businessScore.PeopleNum)
		mp2["score"] = (businessScore.ScoreNum + data.Score) / float64(businessScore.PeopleNum+1)
		err = tx.Model(&model.BusinessScore{}).Where("id = ?", data.BusinessID).Updates(mp1).Error
		if err != nil {
			tx.RollbackTo("test1")
		}
		err = tx.Model(&model.Business{}).Where("id = ?", data.BusinessID).Updates(mp2).Error
		if err != nil {
			tx.RollbackTo("test1")
		}
	}
	return tx.Commit().Error
}
