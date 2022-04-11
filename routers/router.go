package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	v1 "myBlog/api/v1"
	"myBlog/config"
	"myBlog/middleware"
)

// Init 初始化路由
func Init() {
	gin.SetMode(config.ServerConfig.Mode)
	engine := gin.Default()
	engine.Use(middleware.GinRequestLog())
	engine.Use(middleware.Cors())
	routerV1 := engine.Group("api/v1")
	{
		// User 模块路由接口
		routerV1.GET("users", v1.GetUsers)
		routerV1.POST("/login", v1.Login)
		routerV1.POST("user", v1.AddUser)
		// Cate 模块路由接口
		routerV1.GET("categories", v1.GetCate)
		routerV1.GET("category/:id", v1.GetCateDetail)
		routerV1.GET("category/:id/articleList", v1.GetArtByCate)
		// Article 模块路由接口
		routerV1.GET("articles", v1.GetArt)
		routerV1.GET("article/:id", v1.GetArtDetail)
	}
	// 认证路由
	authRouterV1 := engine.Group("api/v1")
	authRouterV1.Use(middleware.JwtToken())
	{
		// User 模块路由接口
		authRouterV1.DELETE("user/:id", v1.DelUser)
		authRouterV1.PUT("user/:id", v1.EditUser)
		// Cate 模块路由接口
		authRouterV1.POST("category", v1.AddCate)
		authRouterV1.DELETE("category/:id", v1.DelCate)
		authRouterV1.PUT("category/:id", v1.EditCate)
		// Article 模块路由接口
		authRouterV1.POST("article", v1.AddArt)
		authRouterV1.DELETE("article/:id", v1.DelArt)
		authRouterV1.PUT("article/:id", v1.EditArt)
		// Upload
		authRouterV1.POST("upload", v1.Upload)
	}

	err := engine.Run(":" + config.ServerConfig.HttpPort)
	if err != nil {
		logrus.Error("start http server error")
	}
}
