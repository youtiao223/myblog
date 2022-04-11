package v1

import (
	"github.com/gin-gonic/gin"
	"myBlog/model"
	"myBlog/utils/errorUtils"
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
		"message": errorUtils.GetErrorMsg(code),
	})
}

// DelCate 删除分类
func DelCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DelCateById(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errorUtils.GetErrorMsg(code),
	})
}

// GetCate 分页查询所有分类
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	cates, count := model.SelectCate(pageNum, pageSize)
	code := errorUtils.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    cates,
		"total":   count,
		"message": errorUtils.GetErrorMsg(code),
	})
}

// GetCateDetail 查询单个文章内容
func GetCateDetail(c *gin.Context) {
	var category model.Category
	CateId, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBind(&category)
	code := model.SelectCateById(uint(CateId), &category)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    category,
		"message": errorUtils.GetErrorMsg(code),
	})
}

// EditCate 编辑分类
func EditCate(c *gin.Context) {
	var cate model.Category
	_ = c.ShouldBind(&cate)
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.UpdateCateById(uint(id), &cate)
	if code == errorUtils.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    cate,
			"message": errorUtils.GetErrorMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errorUtils.GetErrorMsg(code),
		})
	}

}
