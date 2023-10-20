package model

import (
	"gorm.io/gorm"
)

// User 用户结构体，面试用户
type User struct {
	gorm.Model
	UserName  string `gorm:"primarykey;column:userName;type:varchar(50);" json:"username"`
	PassWord  string `gorm:"column:passWord;type:varchar(50);" json:"password"`
	Age       int    `gorm:"column:age;type:int(100);" json:"age"`
	Sex       string `gorm:"column:sex;type:varchar(2);" json:"sex"`
	WxUnionId string `gorm:"column:wxUnionId;type:varchar(50);" json:"wxUnionId"`
	WxOpenId  string `gorm:"column:wxOpenId;type:varchar(50);" json:"wxOpenId"`
}

// TableName 实现接口，使得数据库访问操作可以指定数据表
func (User) TableName() string {
	return "user"
}
