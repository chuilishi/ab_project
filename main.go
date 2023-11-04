package main

import (
	_ "ab_project/docs"
	"ab_project/wechat"
)

func main() {
	wechat.Wechat()
	//data := model.TemplateMessage{
	//	WxOpenId:  "osNMd68bPwCn7FRP-NISWGwg0Ybk",
	//	Name:      "chuilishi",
	//	Msg:       "message",
	//	NowStatus: "成功",
	//	HTTP:      "www.baidu.com",
	//}
	//wechat.SendTemplateMessage(data, wechat.GetAccessToken(true))
}
