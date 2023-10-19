package mysqlDB

import (
	"ab_project/global"
	"ab_project/people"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 初始化数据库连接
func InitGrom() {
	dsn := global.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("无法连接数据库！", err)
		return
	}
	err = db.AutoMigrate(&people.User{})
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
}

func FindUserByUsernamePassword(username, password string) (*people.User, error) {
	ret := new(people.User)
	err := DB.Where("username =? AND password = ? ", username, password).First(ret).Error
	return ret, err
}
