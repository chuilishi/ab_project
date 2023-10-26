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

// FindUserByUsernamePassword 用户实现用户登录功能，返回用户结构体
func FindUserByUsernamePassword(username, password string) (*model.User, error) {

}

// IsUserHave 通过WxOpenId查找用户是否存在
func IsUserHave(WxOpenId string) error {
	user := new(model.User)
	err := DB.Where("wxopenid = ?", WxOpenId).First(user).Error
	return err
}

// RegisterUser 实现用户注册
func RegisterUser(user *model.User) (*model.User, error) {

}
