package dao

import "gorm.io/gorm"

type Dao struct {
	engine *gorm.DB
}

func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		engine: db,
	}
}
