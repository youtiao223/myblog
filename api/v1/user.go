package v1

import (
	"github.com/gin-gonic/gin"
	"myBlog/model"
	"myBlog/utils"
	"myBlog/utils/errors"
	"net/http"
	"strconv"
)

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var user model.User
	// 将请求体绑定到结构体中
	_ = c.ShouldBindJSON(&user)
	// 表单验证  todo 抽离出来
	validateStr, validateCode := utils.ValidateStruct(&user)
	if validateCode == errors.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"status":  validateCode,
			"message": validateStr,
		})
		return
	}
	// model层返回错误码
	code := model.InsertUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errors.GetErrorMsg(code),
	})
}

// DelUser 删除用户
func DelUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DelUserById(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errors.GetErrorMsg(code),
	})
}

// GetUsers 分页查询所有用户
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	users, count := model.SelectUsers(pageNum, pageSize)
	code := errors.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    users,
		"total":   count,
		"message": errors.GetErrorMsg(code),
	})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBind(&user)
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.UpdateUserById(uint(id), &user)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    user,
		"message": errors.GetErrorMsg(code),
	})
}
