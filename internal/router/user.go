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

// Login @Summary 登录
// @Produce json
// @Param code query string true "code"
// @Success 200 {object} swagger.LoginSwagger "成功"
// @Failure 400 {object} errcode.Error "入参错误"
// @Failure 500 {object} errcode.Error "数据库操作失败"
// @Failure 500 {object} errcode.Error "微信API请求错误"
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
