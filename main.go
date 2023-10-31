package main

import (
	_ "ab_project/docs"
	"ab_project/wechat"
	"io"
)

//	func main() {
//		//mysqlDB.InitGrom()
//		//r := router.GetRouter()
//		//r.Run(":8080")
//		//wechat.Wechat()
//	}
func main() {
	////8<<20 即 8*2^20=8M
	//r.MaxMultipartMemory = 8 << 20
	//r.POST("/upload", func(c *gin.Context) {
	//	file, err := c.FormFile("import")
	//	if err != nil {
	//		c.String(500, "上传文件出错")
	//	}
	//	c.SaveUploadedFile(file, file.Filename)
	//	c.String(http.StatusOK, file.Filename+"上传成功")
	//})
	resp := wechat.SendTemplateMessage("osNMd68bPwCn7FRP-NISWGwg0Ybk", "74_hQ9Jrhb4wy_0B5EJJU-T3jbwwkyzNUBpj4IYZbO8mBwCxvUZCuOwbSSZsKpJZ6sYfuks_ph5pjPcoSi2OTXqINS3whx3fuRF3FsJbq9Cj41-uaPbnoF3S2K44R8LDFgABAWOQ")
	//respjson := map[string]interface{}{}
	body, _ := io.ReadAll(resp.Body)
	//json.Unmarshal(body, respjson)
	//for key, value := range respjson {
	//	println(key)
	//	println(value)
	//}
	println(string(body))
}
