package model

type TemplateMessage struct {
	WxOpenId  string `form:"wxOpenId"`  //要发送的openid
	Name      string `form:"name"`      //姓名
	Message   string `form:"message"`   //要发送的信息
	NowStatus string `form:"nowStatus"` //目前录取状态
	HTTP      string `form:"HTTP"`      //要跳转的网址
}
