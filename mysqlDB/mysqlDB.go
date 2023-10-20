package mysqlDB

import (
	"ab_project/global"
	"ab_project/model"
	"errors"
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
		return
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
}

// FindUserByUsernamePassword 用户实现用户登录功能，返回用户结构体
func FindUserByUsernamePassword(username, password string) (*model.User, error) {
	ret := new(model.User)
	err := DB.Where("userName =? AND passWord = ? ", username, password).First(ret).Error
	return ret, err
}

// IsUserHave 通过用户名查找用户是否存在
func IsUserHave(username string) error {
	user := new(model.User)
	err := DB.Where("userName = ?", username).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// RegisterUserByUsername 实现用户注册（仅注册账号与密码）
func RegisterUserByUsername(username, password string) error {
	err := IsUserHave(username)
	if err == nil {
		return errors.New("用户已存在")
	}
	user := new(model.User)
	user.UserName = username
	user.PassWord = password
	dberr := DB.Create(&user)
	fmt.Println(dberr.Error)
	return dberr.Error
}
