package v1

import (
	"github.com/gin-gonic/gin"
	"myBlog/model"
	"myBlog/utils"
	"myBlog/utils/errorUtils"
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
	if validateCode == errorUtils.ERROR {
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
		"message": errorUtils.GetErrorMsg(code),
	})
}

// DelUser 删除用户
func DelUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DelUserById(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errorUtils.GetErrorMsg(code),
	})
}

// GetUsers 分页查询所有用户
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	users, count := model.SelectUsers(pageNum, pageSize)
	code := errorUtils.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    users,
		"total":   count,
		"message": errorUtils.GetErrorMsg(code),
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
		"message": errorUtils.GetErrorMsg(code),
	})
}

// EditProfile 修改管理员个人信息设置
func EditProfile(c *gin.Context) {
	var profile model.Profile
	_ = c.ShouldBind(&profile)
	uid, _ := strconv.Atoi(c.Param("id"))
	code := model.UpdateProfile(uint(uid), &profile)
	if code == errorUtils.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    profile,
			"message": errorUtils.GetErrorMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errorUtils.GetErrorMsg(code),
		})
	}

}

// GetProfile 获取当前管理员个人信息
func GetProfile(c *gin.Context) {
	var profile model.Profile
	uid, _ := strconv.Atoi(c.Param("id"))
	code := model.SelectProfileByUid(uint(uid), &profile)
	if code == errorUtils.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    profile,
			"message": errorUtils.GetErrorMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errorUtils.GetErrorMsg(code),
		})
	}
}
