package middleware

import (
	"eat_box/pkg/app"
	"eat_box/pkg/app/response"
	"eat_box/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Newresponse(
				errcode.UnauthorizedAuthIsEmpty.Code(),
				errcode.UnauthorizedAuthIsEmpty.Msg()).SendResponse(c, errcode.UnauthorizedAuthIsEmpty)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		fmt.Println("parts: ", parts)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Newresponse(
				errcode.UnauthorizedTokenError.Code(),
				errcode.UnauthorizedTokenError.Msg()).SendResponse(c, errcode.UnauthorizedTokenError)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		fmt.Println(parts[1])
		mc, err := app.ParseToken(parts[1])
		if err != nil {
			response.Newresponse(
				errcode.UnauthorizedTokenExpired.Code(),
				errcode.UnauthorizedTokenExpired.Msg()).SendResponse(c, errcode.UnauthorizedTokenExpired)

			c.Abort()
			return
		}
		// 将当前请求的stuid信息保存到请求的上下文c上
		fmt.Println("mc.StuID", mc.ID)
		c.Set("id", mc.ID)
		c.Next() // 后续的处理函数可以用过c.Get("stuid")来获取当前请求的用户信息
	}
}
