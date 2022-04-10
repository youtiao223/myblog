package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {

	return cors.New(cors.Config{

		AllowOrigins: []string{"*"},

		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},

		AllowHeaders: []string{"*"},

		ExposeHeaders: []string{"Content-Length", "Authorization", "Content-Type"},

		// 可选字段是否允许发送cookie,这个值也只能设为true，如果服务器不要浏览器发送Cookie，删除该字段。
		// AllowCredentials: true,

		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},

		MaxAge: 12 * time.Hour,
	})

}
