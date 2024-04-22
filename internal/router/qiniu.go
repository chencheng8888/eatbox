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
func (q Qiniu) SendQNToken(c *gin.Context) {
	svc := service.NewService()
	token := svc.GetQNToken()
	data := swagger.QiniuResponse{
		Token: token,
	}
	response.OK.WithData(data).SendResponse(c, errcode.Success)
}
