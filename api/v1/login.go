package v1

import (
	"github.com/gin-gonic/gin"
	"myBlog/model"
	"myBlog/utils/errors"
	"net/http"
)

// Login 登录
func Login(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)

	token, code := model.CheckLogin(&user)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"token":   token,
		"message": errors.GetErrorMsg(code),
	})

}
