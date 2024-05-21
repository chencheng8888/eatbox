package router

import (
	"eat_box/internal/service"
	"eat_box/pkg/app"
	"eat_box/pkg/app/response"
	"eat_box/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Order struct{}

func NewOrder() Order {
	return Order{}
}

// Create 创建订单
// @Summary 创建订单
// @Produce json
// @Param order formData service.CreateOrderRequest true "订单"
// @Success 200 {object} swagger.CreateOrderResponse "成功"
// @Failure 400 {object} swagger.Fail "入参错误"
// @Failure 500 {object} swagger.Fail "服务端出现问题"
// @Failure 204 {object} swagger.Fail "登录状态有误"
// @Failure 401 {object} swagger.Fail "鉴权失败"
// @Router /api/order/new [post]
func (or Order) Create(c *gin.Context) {
	//参数获取和检验
	param := service.CreateOrderRequest{}
	valid, errs := app.BindAndValid(c, &param)
	if valid {
		fmt.Println(errs) //模拟日志打印

		response.Newresponse(
			errcode.InvalidParams.Code(),
			errcode.InvalidParams.Msg()).SendResponse(c, errcode.InvalidParams)
		return
	}
	svc := service.NewService()
	err := svc.CreateOrder(&param)
	if err.Code() != errcode.Success.Code() {
		response.Newresponse(err.Code(), err.Msg()).SendResponse(c, err)
		return
	}
	response.OK.SendResponse(c, errcode.Success)
}

// Get 获取订单
// @Summary 获取订单
// @Produce json
// @Param id query int true "订单"
// @Success 200 {object} swagger.GetOrderResponse "成功"
// @Failure 400 {object} swagger.Fail "入参错误"
// @Failure 500 {object} swagger.Fail "服务端出现问题"
// @Failure 204 {object} swagger.Fail "登录状态有误"
// @Failure 401 {object} swagger.Fail "鉴权失败"
// @Router /api/order/get [get]
func (or Order) Get(c *gin.Context) {
	//参数获取和检验
	param := service.GetOrderRequest{}
	valid, errs := app.BindAndValid(c, &param)
	if valid {
		fmt.Println(errs) //模拟日志打印

		response.Newresponse(
			errcode.InvalidParams.Code(),
			errcode.InvalidParams.Msg()).SendResponse(c, errcode.InvalidParams)
		return
	}
	svc := service.NewService()
	order, err := svc.GetOrderByID(&param)
	if err.Code() != errcode.Success.Code() {
		response.Newresponse(err.Code(), err.Msg()).SendResponse(c, err)
		return
	}
	response.OK.WithData(order).SendResponse(c, errcode.Success)
}
