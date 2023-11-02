package main

import (
	_ "ab_project/docs"
	"ab_project/model"
	"ab_project/wechat"
)

func main() {
	//r := gin.Default()
	//r.POST("/templateMessage", wechat.TemplateMessageHandler)
	//err := r.Run(":80")
	//if err != nil {
	//	println("错误")
	//}
	//这个是测试,直接发送,上面的是被请求触发之后发送

	data := model.TemplateMessage{
		WxOpenId:  "osNMd68bPwCn7FRP-NISWGwg0Ybk",
		Name:      "chuilishi",
		Message:   "恭喜被录取",
		NowStatus: "成功",
		HTTP:      "www.baidu.com",
	}
	wechat.SendTemplateMessage(data, wechat.GetAccessToken(true))
}
