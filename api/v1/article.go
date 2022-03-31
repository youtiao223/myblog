package v1

import (
	"github.com/gin-gonic/gin"
	"myBlog/model"
	"myBlog/utils/errors"
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
		"message": errors.GetErrorMsg(code),
	})
}

// DelArt 删除文章
func DelArt(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBind(&article)
	code := model.DelArtById(article.ID)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errors.GetErrorMsg(code),
	})
}

// GetArt 获取文章列表
func GetArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	articles := model.SelectArt(pageNum, pageSize)
	code := errors.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    articles,
		"message": errors.GetErrorMsg(code),
	})
}

// EditArt 编辑文章
func EditArt(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBind(&article)
	code := model.UpdateArtById(article.ID, &article)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errors.GetErrorMsg(code),
	})
}

// GetArtDetail 查询单个文章内容
func GetArtDetail(c *gin.Context) {
	var article model.Article
	articleId, _ := strconv.Atoi(c.Query("id"))
	_ = c.ShouldBind(&article)
	code := model.SelectArtById(uint(articleId), &article)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errors.GetErrorMsg(code),
	})
}

// GetArtByCate 查询某个分类下的所有文章
func GetArtByCate(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Query("cid"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	articles, code := model.SelectArtByCate(uint(cid), pageNum, pageSize)
	_ = c.ShouldBind(&articles)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    articles,
		"message": errors.GetErrorMsg(code),
	})
}
