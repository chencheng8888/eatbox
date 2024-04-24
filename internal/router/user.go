package router

import (
	"eat_box/internal/model/swagger"
	"eat_box/internal/service"
	"eat_box/pkg/app"
	"eat_box/pkg/app/response"
	"eat_box/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

// Login
// @Summary 登录
// @Produce json
// @Param code query string true "code"
// @Success 200 {object} swagger.LoginSwagger "成功"
// @Failure 400 {object} swagger.Fail "入参错误"
// @Failure 500 {object} swagger.Fail "数据库操作失败"
// @Failure 500 {object} swagger.Fail "微信API请求错误"
// @Router /api/user/login [post]
func (u User) Login(c *gin.Context) {
	param := service.LoginRequest{}
	valid, errs := app.BindAndValid(c, &param)
	if valid {
		fmt.Println(errs) //模拟日志打印

		response.Newresponse(
			errcode.InvalidParams.Code(),
			errcode.InvalidParams.Msg()).SendResponse(c, errcode.InvalidParams)
		return
	}
	svc := service.NewService()
	first, token, err := svc.Login(&param)
	if err.Code() != errcode.Success.Code() {
		response.Newresponse(err.Code(), err.Msg()).SendResponse(c, err)
		return
	}
	data := swagger.NewLoginData(token, first)
	response.OK.WithData(data).SendResponse(c, errcode.Success)
}

// UpdateUserInfo
// @Summary 更新个人信息
// @Description 这几个参数你可以根据情况来传,不一定非要以forData的形式来传，你也可以传query，根据情况来传就好
// @Produce json
// @Param nickname formData string false "昵称"
// @Param tele formData string false "电话号码"
// @Param headimage formData string false "头像URL"
// @Success 200 {object} swagger.UpdateSwagger "成功"
// @Failure 400 {object} swagger.Fail "入参错误"
// @Failure 500 {object} swagger.Fail "服务端出现问题"
// @Failure 204 {object} swagger.Fail "登录状态有误"
// @Failure 401 {object} swagger.Fail "鉴权失败"
// @Router /api/user/update [put]
func (u User) UpdateUserInfo(c *gin.Context) {
	param := service.UpdateInfoRequest{}
	valid, errs := app.BindAndValid(c, &param)
	if valid {
		fmt.Println(errs) //模拟日志打印

		response.Newresponse(
			errcode.InvalidParams.Code(),
			errcode.InvalidParams.Msg()).SendResponse(c, errcode.InvalidParams)
		return
	}
	id, _ := app.GetIDFromToken(c)
	param.ID = id
	svc := service.NewService()
	err := svc.UpdateInfo(&param)
	if err.Code() != errcode.Success.Code() {
		response.Newresponse(err.Code(), err.Msg()).SendResponse(c, err)
		return
	}
	response.OK.SendResponse(c, errcode.Success)
}

// GetUserInfo
// @Summary 获取用户信息
// @Description 这个ID不是必须的，也就是说，如果你不填默认就是查自己的。而是不是自己的，返回的信息也有差别
// @Produce json
// @Param id query string false "用户ID"
// @Success 200 {object} swagger.DetailSwagger "成功"
// @Failure 400 {object} swagger.Fail "入参错误"
// @Failure 500 {object} swagger.Fail "服务端出现错误"
// @Failure 204 {object} swagger.Fail "登录状态有误"
// @Failure 401 {object} swagger.Fail "鉴权失败"
// @Router /api/user/getinfo [get]
func (u User) GetUserInfo(c *gin.Context) {
	param := service.DetailRequest{}
	valid, errs := app.BindAndValid(c, &param)
	if valid {
		fmt.Println(errs) //模拟日志打印

		response.Newresponse(
			errcode.InvalidParams.Code(),
			errcode.InvalidParams.Msg()).SendResponse(c, errcode.InvalidParams)
		return
	}
	id, _ := app.GetIDFromToken(c)
	if param.ID == "" {
		param.ID = id
		param.Self = true
	} else {
		param.Self = id == param.ID
	}
	svc := service.NewService()
	data, err := svc.GetUserInfo(&param)
	if err.Code() != errcode.Success.Code() {
		response.Newresponse(err.Code(), err.Msg()).SendResponse(c, err)
		return
	}
	response.OK.WithData(data).SendResponse(c, errcode.Success)
}

// GetBusinessList
// @Summary 获取商家列表
// @Produce json
// @Param page query string true "页码"
// @Success 200 {object} swagger.BusinessListSwagger "成功"
// @Failure 400 {object} swagger.Fail "入参错误或者页码超限"
// @Failure 500 {object} swagger.Fail "服务端出现错误"
// @Failure 204 {object} swagger.Fail "登录状态有误"
// @Failure 401 {object} swagger.Fail "鉴权失败"
// @Router /api/user/getbusinesslist [get]
func (u User) GetBusinessList(c *gin.Context) {
	param := service.GetBusinessesRequest{}
	valid, errs := app.BindAndValid(c, &param)
	if valid {
		fmt.Println(errs) //模拟日志打印

		response.Newresponse(
			errcode.InvalidParams.Code(),
			errcode.InvalidParams.Msg()).SendResponse(c, errcode.InvalidParams)
		return
	}
	svc := service.NewService()
	data, err := svc.GetBusinesses(&param)
	if err.Code() != errcode.Success.Code() {
		response.Newresponse(err.Code(), err.Msg()).SendResponse(c, err)
		return
	}
	response.OK.WithData(data).SendResponse(c, errcode.Success)
}
