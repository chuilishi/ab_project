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

	err = c.SaveUploadedFile(file, "")
	if err != nil {
		return
	}
	c.String(http.StatusOK, file.Filename+"上传成功")
}
