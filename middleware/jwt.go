package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"myBlog/config"
	"myBlog/utils/errors"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte(config.ServerConfig.JwtKey)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成token
func GenToken(username string) (string, int) {
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
func ValidateToken(tokenString string) (*MyCustomClaims, bool) {
	token, _ := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取http请求头中的认证信息
		tokenHeader := c.Request.Header.Get("Authorization")
		code := errors.SUCCESS

		// 请求头中没有认证信息
		if tokenHeader == "" {
			code = errors.ErrorTokenExits
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errors.GetErrorMsg(code),
			})
			c.Abort()
			return
		}

		checkedToken := strings.SplitN(tokenHeader, " ", 2)

		// 检查Authorization头格式
		// Authorization: <type> <credentials>
		// 格式错误
		if len(checkedToken) != 2 || checkedToken[0] != "Bearer" {
			code = errors.ErrorTokenFormat
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errors.GetErrorMsg(code),
			})
			c.Abort()
			return
		}

		// 验证Token签名
		key, ok := ValidateToken(checkedToken[1])
		// Token 验证错误
		if ok == false {
			code = errors.ErrorTokenValidate
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errors.GetErrorMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Next()
	}
}
