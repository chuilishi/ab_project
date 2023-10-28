package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传文件出错")
	}
	c.SaveUploadedFile(file, file.Filename)
	c.String(http.StatusOK, file.Filename+"上传成功")
}
