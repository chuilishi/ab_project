package model

type TemplateMessage struct {
	WxOpenId  string `form:"wxOpenId" json:"wxopenid"`   //要发送的openid
	Name      string `form:"name" json:"name"`           //姓名
	Msg       string `form:"msg" json:"msg"`             //要发送的信息
	NowStatus string `form:"nowStatus" json:"nowstatus"` //目前录取状态
	HTTP      string `form:"HTTP" json:"HTTP"`           //要跳转的网址
	Code      int    `form:"code" json:"code"`           //发送消息的标题类型
}
