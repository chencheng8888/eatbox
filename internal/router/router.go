package router

import (
	"eat_box/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	QiniuGroup := r.Group("/qiniu")
	{
		qn := NewQiniu()
		QiniuGroup.GET("/get_token", middleware.JWTAuthMiddleware(), qn.SendQNToken)
	}
	UserGroup := r.Group("/api/user")
	{
		user := NewUser()
		UserGroup.POST("/login", user.Login)
		UserGroup.PUT("/update", middleware.JWTAuthMiddleware(), user.UpdateUserInfo)
		UserGroup.GET("/getinfo", middleware.JWTAuthMiddleware(), user.GetUserInfo)
		UserGroup.GET("/getbusinesslist", middleware.JWTAuthMiddleware(), user.GetBusinessList)
	}
	return r
}
