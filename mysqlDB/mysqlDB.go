package mysqlDB

import (
	"ab_project/global"
	"ab_project/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitGrom 初始化数据库连接
func InitGrom() {
	dsn := global.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("无法连接数据库！", err)
		panic("无法连接数据库！")
		return
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.AutoMigrate(&model.Manager{})
	if err != nil {
		fmt.Println(err)
		return
	}

	DB = db
}

// IsUserHave 通过WxOpenId查找用户是否存在
func IsUserHave(WxOpenId string) *model.User {
	user := new(model.User)
	DB.Where("wxopenid = ?", WxOpenId).First(user)
	return user
}

// RegisterUser 实现用户简历投递
func RegisterUser(user *model.User) error {
	tempuser := IsUserHave(user.WxOpenId)
	if tempuser.ID == 0 {

		return DB.Create(user).Error
	}
	return DB.Updates(user).Error
}

// 实现返回指定方向用户信息
func FindUsersByDirection(direction string) []model.User {
	var users []model.User
	if direction == "全部" {
		DB.Find(&users)

	} else {
		DB.Where("direction = ?", direction).Find(&users)
	}
	return users
}
