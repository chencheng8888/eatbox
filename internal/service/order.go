package service

import (
	"eat_box/internal/model"
	"eat_box/pkg/errcode"
)

type CreateOrderRequest struct {
	Num          int    `form:"num" binding:"required"`           //购买数量
	Status       int    `form:"status" binding:"required"`        //订单状态
	PayStatus    int    `form:"pay_status" binding:"required"`    //支付状态
	FinishedTime string `form:"finished_time" binding:"required"` //完成时间
	Pay          int    `form:"pay" binding:"required"`           //支付金额
	Addr         string `form:"addr" binding:"required"`          //配送地址
	AddrWay      string `form:"addr_way" binding:"required"`      //配送方式
	BoxID        uint   `form:"box_id" binding:"required"`        //盲盒ID
	UserID       string `form:"user_id" binding:"required"`       //用户ID
	BusID        uint   `form:"bus_id" binding:"required"`        //商家ID
}
type GetOrderRequest struct {
	ID uint `form:"id" binding:"required"` //订单ID
}

func (svc *Service) CreateOrder(param *CreateOrderRequest) *errcode.Error {
	order := model.NewOrder(
		param.Status,
		param.PayStatus,
		param.Pay,
		param.Num,
		param.FinishedTime,
		param.Addr,
		param.AddrWay,
		param.BoxID,
		param.BusID,
		param.UserID,
	)
	err := svc.dao.CreateOrder(order)
	if err != nil {
		return errcode.MySQLErr
	}
	return errcode.Success
}
func (svc *Service) GetOrderByID(param *GetOrderRequest) (model.Order, *errcode.Error) {
	order, err := svc.dao.FindOrderByID(param.ID)
	if err != nil {
		return model.Order{}, errcode.MySQLErr
	}
	return order, errcode.Success
}
