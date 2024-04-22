package app

import (
	"eat_box/global"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

// 生成jwt
func GenerateToken(ID string) (string, error) {
	c := Claims{
		ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(global.TokenExpireDuration).Unix(),
			Issuer:    "my-project",
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(global.Secret)
}

// 解析JWT
func ParseToken(tokenString string) (*Claims, error) {
	//解析Token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return global.Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
func GetIDFromToken(c *gin.Context) (string, bool) {
	id, ok := c.Get("id")
	return id.(string), ok
}
