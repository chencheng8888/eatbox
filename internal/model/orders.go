package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Num          int    `json:"num" gorm:"column:num"`                     //购买数量
	Status       int    `json:"status" gorm:"column:status"`               //订单状态
	PayStatus    int    `json:"pay_status" gorm:"column:pay_status"`       //支付状态
	FinishedTime string `json:"finished_time" gorm:"column:finished_time"` //完成时间
	Pay          int    `json:"pay" gorm:"column:pay"`                     //支付金额
	Addr         string `json:"addr" gorm:"column:addr"`                   //配送地址
	AddrWay      string `json:"addr_way" gorm:"column:addr_way"`           //配送方式
	BoxID        uint   `json:"box_id" gorm:"column:box_id"`               //盲盒ID
	UserID       string `json:"user_id" gorm:"column:user_id"`             //用户ID
	BusID        uint   `json:"bus_id" gorm:"column:bus_id"`               //商家ID
}
type OrderAndUserID struct {
	gorm.Model
	UserID  string `json:"user_id"`
	OrderID uint   `json:"order_id"`
}
type OrderAndBusID struct {
	gorm.Model
	BusID   uint `json:"bus_id"`
	OrderID uint `json:"order_id"`
}

func NewOrder(sta, paysta, price, num int, ft, addr, addrway string, boxid, busid uint, userid string) Order {
	return Order{
		Num:          num,
		Status:       sta,
		PayStatus:    paysta,
		Pay:          price,
		FinishedTime: ft,
		Addr:         addr,
		AddrWay:      addrway,
		BoxID:        boxid,
		BusID:        busid,
		UserID:       userid,
	}
}
func NewOAU(userid string, orderid uint) OrderAndUserID {
	return OrderAndUserID{
		UserID:  userid,
		OrderID: orderid,
	}
}
func NewOAB(busid, orderid uint) OrderAndBusID {
	return OrderAndBusID{
		BusID:   busid,
		OrderID: orderid,
	}
}
