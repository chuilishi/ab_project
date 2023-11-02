package main

import (
	_ "ab_project/docs"
	"ab_project/wechat"
	"github.com/gin-gonic/gin"
)

//	func main() {
//		//mysqlDB.InitGrom()
//		//r := router.GetRouter()
//		//r.Run(":8080")
//		//wechat.Wechat()
//	}
func main() {
	r := gin.Default()
	r.POST("/templateMessage", wechat.TemplateMessageHandler)
	err := r.Run(":80")
	if err != nil {
		println("错误")
	}

	//这个是测试,直接发送,上面的是被请求触发之后发送

	//data := model.TemplateMessage{
	//	WxOpenId:  "osNMd68bPwCn7FRP-NISWGwg0Ybk",
	//	Name:      "chuilishi",
	//	Message:   "测试",
	//	NowStatus: "成功录取",
	//	HTTP:      "www.baidu.com",
	//}
	//wechat.SendTemplateMessage(data, wechat.GetAccessToken(true))
}
