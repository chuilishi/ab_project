package model

type TemplateMessage struct {
	WxOpenId  string `json:"wxOpenId"`  //要发送的openid
	Name      string `json:"name"`      //姓名
	Message   string `json:"message"`   //要发送的信息
	NowStatus string `json:"nowStatus"` //目前录取状态
	HTTP      string `json:"HTTP"`      //要跳转的网址
}
