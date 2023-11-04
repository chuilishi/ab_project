package model

type UserFile struct {
	WxOpenid string `gorm:"primarykey;column:wxopenid;type:varchar(50);" json:"wxopenid"` //微信openid
	FileName string `gorm:"filename;column:filename;type:varchar(50);" json:"filename"`   //用户文件名称
}

func (UserFile) TableName() string {
	return "sys_userfile"
}
