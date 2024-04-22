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
	QiniuGroup := r.Group("/qiniu")
	{
		qn := NewQiniu()
		QiniuGroup.GET("/get_token", qn.SendQNToken)
	}
	UserGroup := r.Group("/api/user")
	{
		user := NewUser()
		UserGroup.POST("/login", user.Login)
		UserGroup.PUT("/update", user.UpdateUserInfo)
	}
	return r
}
