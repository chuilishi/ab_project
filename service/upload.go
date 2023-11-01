package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传文件出错")
	}
	dst := path.Join("./statics", file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.String(http.StatusOK, file.Filename+"上传成功")
}
