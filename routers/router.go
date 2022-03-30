package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	v1 "myBlog/api/v1"
	"myBlog/config"
)

// Init 初始化路由
func Init() {
	gin.SetMode(config.ServerConfig.Mode)
	engine := gin.Default()

	routerV1 := engine.Group("api/v1")
	{
		// User 模块路由接口
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("user/getUsers", v1.GetUsers)
		routerV1.DELETE("user/del", v1.DelUser)
		routerV1.PUT("user/edit", v1.EditUser)
		// Cate 模块路由接口
		routerV1.POST("cate/add", v1.AddCate)
		routerV1.GET("cate/getCates", v1.GetCate)
		routerV1.DELETE("cate/del", v1.DelCate)
		routerV1.PUT("cate/edit", v1.EditCate)
	}

	err := engine.Run(":" + config.ServerConfig.HttpPort)
	if err != nil {
		logrus.Error("start http server error")
	}
}
