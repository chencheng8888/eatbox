package router

import (
	"eat_box/internal/service"
	"eat_box/pkg/app"
	"eat_box/pkg/app/response"
	"eat_box/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Business struct{}

func NewBusiness() Business {
	return Business{}
}

func (b Business) Create(c *gin.Context) {
	//参数获取和检验
	param := service.CreateBusinessRequest{}
	valid, errs := app.BindAndValid(c, &param)
	if valid {
		fmt.Println(errs) //模拟日志打印

		response.Newresponse(
			errcode.InvalidParams.Code(),
			errcode.InvalidParams.Msg()).SendResponse(c, errcode.InvalidParams)
		return
	}
	svc := service.NewService()
	err := svc.CreateBusiness(&param)
	if err.Code() != errcode.Success.Code() {
		response.Newresponse(err.Code(), err.Msg()).SendResponse(c, err)
		return
	}
	response.OK.SendResponse(c, errcode.Success)
}
