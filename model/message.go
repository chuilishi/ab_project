package model

type WeChatMessage struct {
	WxOpenId  string //要发送的openid
	Name      string //姓名
	Message   string //要发送的信息
	NowStatus string //目前录取状态
	HTTP      string //要跳转的网址
}
