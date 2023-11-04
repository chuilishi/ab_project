package mysqlDB

import (
	"ab_project/global"
	"ab_project/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
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
	err = os.Mkdir("./userFile", 0755)
	DB = db
	if err != nil {
		fmt.Println("无法创建userfile文件夹")
		return
	}

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
		err := os.Mkdir("./userFile/"+user.WxOpenId, 0755)
		if err != nil {
			return err
		}
		err = DB.Create(user).Error
		if err != nil {
			fmt.Println("创建用户数据出错" + err.Error())
		}
		user.PersonalId = fmt.Sprintf("%d%02d%04d", time.Now().Year(), time.Now().Month(), IsUserHave(user.WxOpenId).ID)
	}

	return DB.Updates(user).Error

}

// FindUsersByDirection 实现返回指定方向用户信息
func FindUsersByDirection(direction string) []model.User {
	var users []model.User
	if direction == "全部" {
		DB.Find(&users)

	} else {
		DB.Where("direction = ?", direction).Find(&users)
	}
	return users
}

// FindUsersByStatus 实现返回指定状态用户信息
func FindUsersByStatus(direction string) []model.User {
	var users []model.User

	DB.Where("status = ?", direction).Find(&users)

	return users
}
