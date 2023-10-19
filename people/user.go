package people

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName  string `gorm:"primarykey;column:username;type:varchar(50);" json:"username"`
	PassWord  string `gorm:"column:password;type:varchar(50);" json:"password"`
	Age       int    `gorm:"column:age;type:int(100);" json:"age"`
	Sex       string `gorm:"column:sex;type:varchar(2);" json:"sex"`
	WxUnionId string `gorm:"column:wxUnionId;type:varchar(50);" json:"wxUnionId"`
	WxOpenId  string `gorm:"column:wxOpenId;type:varchar(50);" json:"wxOpenId"`
}
type LoginUser struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

// 实现接口，使得数据库访问操作可以指定数据表
func (User) TableName() string {
	return "users"
}
