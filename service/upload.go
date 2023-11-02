package service

import (
	"ab_project/service/response"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传文件出错")
	}
	if file.Size > 1024*1024*8 {
		response.FailWithMessage("文件超出大小限制", c)
		return
	}
	err = c.SaveUploadedFile(file, "")
	if err != nil {
		return
	}
	response.OkWithMessage("上传成功", c)
}
