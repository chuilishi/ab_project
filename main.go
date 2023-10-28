package main

import (
	_ "ab_project/docs"
	"github.com/gin-gonic/gin"
	"net/http"
)

//	func main() {
//		//mysqlDB.InitGrom()
//		//r := router.GetRouter()
//		//r.Run(":8080")
//		//wechat.Wechat()
//	}
func main() {
	r := gin.Default()
	//8<<20 即 8*2^20=8M
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("import")
		if err != nil {
			c.String(500, "上传文件出错")
		}
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename+"上传成功")
	})
	r.Run()
}
