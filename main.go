package main

import (
	_ "ab_project/docs"
	"ab_project/model"
	"ab_project/wechat"
)

//	func main() {
//		//mysqlDB.InitGrom()
//		//r := router.GetRouter()
//		//r.Run(":8080")
//		//wechat.Wechat()
//	}
func main() {
	//r := gin.Default()
	//r.POST("/templateMessage", wechat.TemplateMessageHandler)
	//err := r.Run(":80")
	//if err != nil {
	//	println("错误")
	//}
	data := model.TemplateMessage{
		WxOpenId:  "osNMd6yQN02kDAW7UiNsotP8J1YU",
		Name:      "chuilishi",
		Message:   "测试",
		NowStatus: "成功录取",
		HTTP:      "www.baidu.com",
	}
	wechat.SendTemplateMessage(data, "74_iI5-wV-sik6z_qrM9YQdu5RWxtgRweHjxuoHSQcfjBbIhCWU9Ko8ZVNArG3zyecEMlZmlzqr_Nq5k2NDa3vHVwqLD12_Yh8JjjttH7dGQmLT2M60JQE2iQZFx28VNTeAIAYQW")
}
