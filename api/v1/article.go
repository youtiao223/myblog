package v1

import (
	"github.com/gin-gonic/gin"
	"myBlog/model"
	"myBlog/utils/errorUtils"
	"net/http"
	"strconv"
)

// AddArt 添加文章
func AddArt(c *gin.Context) {
	var article model.Article
	// 将请求体绑定到结构体中
	_ = c.ShouldBindJSON(&article)

	// model层返回错误码
	code := model.InsertArt(&article)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errorUtils.GetErrorMsg(code),
	})
}

// DelArt 删除文章
func DelArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DelArtById(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errorUtils.GetErrorMsg(code),
	})
}

// GetArt 获取文章列表
func GetArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	articles, count := model.SelectArt(pageNum, pageSize)
	code := errorUtils.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    articles,
		"total":   count,
		"message": errorUtils.GetErrorMsg(code),
	})
}

// EditArt 编辑文章
func EditArt(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBind(&article)
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.UpdateArtById(uint(id), &article)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errorUtils.GetErrorMsg(code),
	})
}

// GetArtDetail 查询单个文章内容
func GetArtDetail(c *gin.Context) {
	var article model.Article
	articleId, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBind(&article)
	code := model.SelectArtById(uint(articleId), &article)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errorUtils.GetErrorMsg(code),
	})
}

// GetArtByCate 查询某个分类下的所有文章
func GetArtByCate(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	articles, code := model.SelectArtByCate(uint(cid), pageNum, pageSize)
	_ = c.ShouldBind(&articles)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    articles,
		"message": errorUtils.GetErrorMsg(code),
	})
}
