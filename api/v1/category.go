package v1

import (
	"github.com/gin-gonic/gin"
	"myBlog/model"
	"myBlog/utils/errors"
	"net/http"
	"strconv"
)

// AddCate 添加分类
func AddCate(c *gin.Context) {
	var cate model.Category
	// 将请求体绑定到结构体中
	_ = c.ShouldBindJSON(&cate)

	// model层返回错误码
	code := model.InsertCate(&cate)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errors.GetErrorMsg(code),
	})
}

// DelCate 删除分类
func DelCate(c *gin.Context) {
	var cate model.Category
	_ = c.ShouldBind(&cate)
	code := model.DelCateById(cate.Id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errors.GetErrorMsg(code),
	})
}

// GetCate 分页查询所有分类
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	cates := model.SelectCate(pageNum, pageSize)
	code := errors.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    cates,
		"message": errors.GetErrorMsg(code),
	})
}

// EditCate 编辑分类
func EditCate(c *gin.Context) {
	var cate model.Category
	_ = c.ShouldBind(&cate)
	code := model.UpdateCateById(cate.Id, &cate)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    cate,
		"message": errors.GetErrorMsg(code),
	})
}
