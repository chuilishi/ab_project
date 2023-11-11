package mysqlDB

import (
	"ab_project/global"
	"ab_project/model"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
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
	err = db.AutoMigrate(&model.User{}, &model.Message{}, model.Manager{})
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

// IsUserHaveByID 通过WxOpenId查找用户是否存在
func IsUserHaveByID(id string) *model.User {
	user := new(model.User)
	DB.Where("id = ?", id).First(user)
	return user
}

func IsUserHaveByStudentId(studentId string) *model.User {
	user := new(model.User)
	DB.Where("studentid = ?", studentId).Find(user)
	return user
}

// RegisterUser 实现用户简历投递
func RegisterUser(user *model.User) error {
	tempuser := IsUserHaveByStudentId(user.StudentId)
	if tempuser.ID == 0 {
		err := DB.Create(user).Error
		if err != nil {
			fmt.Println("创建用户数据出错" + err.Error())
		}
		user = IsUserHave(user.WxOpenId)

		err = os.Mkdir("./userFile/"+strconv.Itoa(int(user.ID)), 0755)
		return nil
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

// FindProblemUsers 实现返回异常状态用户
func FindProblemUsers() []model.User {
	var users []model.User
	DB.Where("isproblem = 1").Find(&users)
	return users
}

// SendMessageToUser 通过wxopenid来记录给用户发的消息
func SendMessageToUser(wxopenid string, message string, code int) error {
	user := IsUserHave(wxopenid)
	if user.ID == 0 {
		return errors.New("无法通过wxopenid找到用户")
	}
	var msg = model.Message{
		Message: message,
		UserID:  user.ID,
		Code:    code,
	}
	return DB.Create(&msg).Error
}

// FindUserMessage 通过wxopenid查找给用户发送的消息
func FindUserMessage(wxopenid string) ([]model.Message, error) {
	user := IsUserHave(wxopenid)
	if user.ID == 0 {
		return nil, errors.New("无法通过wxopenid找到用户")
	}
	DB.Where("wxopenid = ?", wxopenid).Preload("Messages").Preload("User").Take(&user)
	return user.Messages, nil

}

// ChangeUserStatus 修改用户状态
func ChangeUserStatus(userid int, status string) error {
	user := IsUserHaveByID(strconv.Itoa(userid))
	if user.ID == 0 {
		return errors.New("无法查找到学号" + strconv.Itoa(userid))
	}
	user.Status = status
	return DB.Updates(user).Error
}

// PostProblem 用户提交问题
func PostProblem(wxopenid string, problem string) error {
	user := IsUserHave(wxopenid)
	if user.ID == 0 {
		return errors.New("无法查找到wxopenid")
	}
	user.ISProblem = 1
	user.Problem = problem
	return DB.Updates(user).Error
}
func AllUserStatus() (interface{}, error) {
	var user []struct {
		Id     uint   `gorm:"id" json:"id"`
		Status string `gorm:"status" json:"status"`
	}
	err := DB.Model(&model.User{}).Select("id,status").Find(&user).Error
	return user, err
}
