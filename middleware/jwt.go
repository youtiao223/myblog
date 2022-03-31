package middleware

import (
	"github.com/golang-jwt/jwt"
	"myBlog/config"
	"myBlog/utils/errors"
	"time"
)

var jwtKey = []byte(config.ServerConfig.JwtKey)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成token
func GenToken(username string, password string) (string, int) {
	// token 过期时间
	expireTime := time.Now().Add(10 * time.Hour)

	// Create the Claims
	claims := MyCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "zyw",
		},
	}

	// 签名算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.ERROR
	}

	return ss, errors.SUCCESS
}

// ValidateToken 验证token
func ValidateToken(tokenString string) (*MyCustomClaims, int) {
	token, _ := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, errors.SUCCESS
	} else {
		return nil, errors.ERROR
	}
}

// todo JwtToken jwt中间件
//func JwtToken() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		tokenHeader := c.Request.Header.Get("Authorization")
//		code := errors.SUCCESS
//		if tokenHeader == "" {
//			code = errors.
//		}
//	}
//}
