package main

import (
	_ "ab_project/docs"
	_ "ab_project/global"
	"ab_project/mysqlDB"
	"ab_project/router"
	_ "ab_project/zap"
)

//	func main() {
//		//
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

	mysqlDB.InitGrom()
	go router.InitGetUserFileMessageHandler()

	r := router.GetRouter()
	r.Run(":8080")
	//wechat.Wechat()
	//resp := wechat.SendTemplateMessage("osNMd68bPwCn7FRP-NISWGwg0Ybk", "74_djnRi8NLRpkr8uXd3OPJmNMZ7XAANWTnsw5QBKMZKYl4iuoX1rrYZTmrBy90aNw1QDjWU_2O2aj3oVflruCs_wBBSMcL3_CwuclmMXk_It0BwINHUMy7ftqeNOsRWBaAGAFKP")
	//body, _ := io.ReadAll(resp.Body)
	//println(string(body))

}
