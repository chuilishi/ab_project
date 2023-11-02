package service

import (
	"ab_project/service/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("上传文件出错", c)
	}
	dst := path.Join("./statics", "Picture.jpg")
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		fmt.Println(err)
		return
	}
	response.OkWithMessage("上传文件成功", c)
}
