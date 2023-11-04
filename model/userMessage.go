package model

type UserMessage struct {
	Model
	WxOpenId string `gorm:"primarykey;column:wxopenid;type:varchar(50);default:'未知';" json:"wxopenid"` //微信openid
	message  string `gorm:"column:message"`
}
