package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"myBlog/utils"
	"time"
)

// GinRequestLog gin api 请求纪录( 请求状态码 处理时间 请求方法 IP 路由 )
func GinRequestLog() gin.HandlerFunc {
	var ginLogger = logrus.New()
	ginLogger.SetLevel(logrus.DebugLevel)

	// 可以指定存放日志的目录,默认为/logs,支持指定多个目录
	// 例如若想更改目录为/log2
	// 调用 WriteLogToFolder(ginLogger, "log2")
	utils.WriteLogToFolder(ginLogger)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		ginEntry := ginLogger.WithFields(logrus.Fields{
			"startTime":   startTime,
			"endTime":     endTime,
			"latencyTime": latencyTime,
			"reqMethod":   reqMethod,
			"reqUri":      reqUri,
			"statusCode":  statusCode,
			"clientIP":    clientIP,
		})

		if len(c.Errors) > 0 {
			ginEntry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			ginEntry.Error()
		} else if statusCode >= 400 {
			ginEntry.Warn()
		} else {
			ginEntry.Info()
		}
	}
}
