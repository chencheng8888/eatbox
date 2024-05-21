package service

import (
	"eat_box/internal/model"
	"eat_box/pkg/errcode"
)

type CreateBusinessRequest struct {
	ManagerID   string `form:"manager_id" binding:"required"`
	Name        string `form:"name" binding:"required"`
	Avatar      string `form:"avatar" binding:"required"`
	Address     string `form:"address" binding:"required"`
	Openhours   string `form:"openhours" binding:"required"`
	Description string `form:"description" binding:"omitempty"`
	Image       string `form:"image" binding:"omitempty"`
	Tele        string `form:"tele" binding:"required"`
}

func (svc *Service) CreateBusiness(param *CreateBusinessRequest) *errcode.Error {
	business := model.Business{
		ManagerID:    param.ManagerID,
		Name:         param.Name,
		Avatar:       param.Avatar,
		Address:      param.Address,
		Openinghours: param.Openhours,
		Description:  param.Description,
		Image:        param.Image,
		Tele:         param.Tele,
	}
	err := svc.dao.CreateBusiness(business)
	if err != nil {
		return errcode.MySQLErr
	}
	return errcode.Success
}
