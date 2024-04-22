package response

import (
	"eat_box/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	OK = Newresponse(errcode.Success.Code(), errcode.Success.Msg())
)

// 自定义响应信息
func (res *Response) WithMsg(message string) Response {
	return Response{
		Code: res.Code,
		Msg:  message,
		Data: res.Data,
	}
}

// 追加响应数据
func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

// 构造函数
func Newresponse(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

// 发送响应
func (res Response) SendResponse(c *gin.Context, err *errcode.Error) {
	c.JSON(err.StatusCode(), res)
}
