package v1

import (
	"github.com/gin-gonic/gin"
	"myBlog/model"
	"myBlog/utils/errorUtils"
	"net/http"
)

// Login 登录
func Login(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)

	token, code := model.CheckLogin(&user)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"token":   token,
		"message": errorUtils.GetErrorMsg(code),
	})

}
