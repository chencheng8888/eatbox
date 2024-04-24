package router

import (
	"eat_box/internal/model/swagger"
	"eat_box/internal/service"
	"eat_box/pkg/app/response"
	"eat_box/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Qiniu struct{}

func NewQiniu() Qiniu {
	return Qiniu{}
}

// SendQNToken
// @Summary 获取七牛的token
// @Produce json
// @Success 200 {object} swagger.QiniuSwagger "成功"
// @Failure 500 {object} swagger.Fail "服务端出现问题"
// @Router /qiniu/get_token [get]
func (q Qiniu) SendQNToken(c *gin.Context) {
	svc := service.NewService()
	token := svc.GetQNToken()
	data := swagger.QiniuData{
		Token: token,
	}
	response.OK.WithData(data).SendResponse(c, errcode.Success)
}
