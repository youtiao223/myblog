package v1

import (
	"github.com/gin-gonic/gin"
	"myBlog/service"
	"myBlog/utils/errors"
	"net/http"
)

// Upload 文件上传
func Upload(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	fileUrl, code := service.UploadFile(file, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"fileUrl": fileUrl,
		"message": errors.GetErrorMsg(code),
	})

}
