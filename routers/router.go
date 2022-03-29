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
		//todo 删除用户,暂时不支持
		//routerV1.POST("user/del", v1.DelUser)

		//
	}

	err := engine.Run(":" + config.ServerConfig.HttpPort)
	if err != nil {
		logrus.Error("start http server error")
	}
}
