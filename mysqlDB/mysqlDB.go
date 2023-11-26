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
	err = db.AutoMigrate(&model.User{}, &model.Message{}, &model.MessageTemplate{})
	if err != nil {
		fmt.Println(err)
		return
	}

	DB = db
	_, err = os.Stat("./userFile")
	err = os.Mkdir("./userFile", 0755)
	DB = db
	if err != nil {
		fmt.Println("无法创建userfile文件夹")
	}
	_, err = os.Stat("./userFile/picture")
	err = os.Mkdir("./userFile/picture", 0755)

	if err != nil {
		fmt.Println("无法创建picture文件夹")
		return
	}

	_, err = os.Stat("./userFile/ppp")
	err = os.Mkdir("./userFile/ppp", 0755)

	if err != nil {
		fmt.Println("无法创建ppp文件夹")
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

// IsUserHaveByStudentId 通过id查抄用户是否存在
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

		return err
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
func SendMessageToUser(wxopenid string, message string, code int, messageid string) error {
	user := IsUserHave(wxopenid)
	if user.ID == 0 {
		return errors.New("无法通过wxopenid找到用户")
	}
	var msg = model.Message{
		Message:   message,
		UserID:    user.ID,
		Code:      code,
		MessageId: messageid,
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

// AllUserStatus 返回所有状态
func AllUserStatus() (interface{}, error) {
	var user []struct {
		Id     uint   `gorm:"id" json:"id"`
		Status string `gorm:"status" json:"status"`
	}
	err := DB.Model(&model.User{}).Select("id,status").Find(&user).Error
	return user, err
}

// FindUserPassHistory 查找不同code的message
func FindUserPassHistory(wxopenid string, code []int) ([]model.Message, error) {
	user := IsUserHave(wxopenid)
	if user.ID == 0 {
		return nil, errors.New("无法通过wxopenid找到用户")
	}
	if code[0] != -1 {
		err := DB.Where("wxopenid = ?", wxopenid).Preload("Messages", "code in ?", code).Take(&user).Error

		return user.Messages, err
	} else {
		err := DB.Where("wxopenid = ?", wxopenid).Preload("Messages").Take(&user).Error
		return user.Messages, err
	}
}

// FindUserPassHistory 查找不同code的message
func FindUserPassHistoryByStudentid(studentid string, code int) ([]model.Message, error) {
	user := IsUserHaveByStudentId(studentid)
	if user.ID == 0 {
		return nil, errors.New("无法通过studentid找到用户")
	}
	if code != -1 {
		err := DB.Where("studentid = ?", studentid).Preload("Messages", "code= ?", code).Take(&user).Error

		return user.Messages, err
	} else {
		err := DB.Where("studentid = ?", studentid).Preload("Messages").Take(&user).Error
		return user.Messages, err

	}
}

// OverComeProblem 处理异常
func OverComeProblem(user *model.User) error {
	user.Problem = "无异常信息"

	err := DB.Updates(user).Error
	DB.Model(&user).Update("isproblem", "0")
	return err

}

// UpdateUser 更新用户
func UpdateUser(user *model.User) error {
	return DB.Updates(user).Error
}

// IsIdDUsed 查询id是否用过
func IsIdDUsed(id string) bool {
	var temp []model.Message
	DB.Model(&model.Message{}).Where("messageid  =? ", id).Find(&temp)
	if len(temp) < 2 {
		return false
	} else {
		return true
	}

}

// IsIDSend 查询id是否发送
func IsIDSend(id string) bool {
	var temp []model.Message
	DB.Model(&model.Message{}).Where("messageid  =? ", id).Find(&temp)
	if len(temp) > 0 {
		return true
	} else {
		return false
	}
}

// SaveMessageTemplate 存入数据库 消息模版
func SaveMessageTemplate(message *model.MessageTemplate) error {
	return DB.Model(&model.MessageTemplate{}).Create(message).Error
}
func DeleteMessageTemplate(id int) error {
	message := new(model.MessageTemplate)
	DB.Model(&model.MessageTemplate{}).Where("id = ?", id).First(message)
	if message.ID == 0 {
		return errors.New("无法查找到消息")
	}

	return DB.Model(&model.MessageTemplate{}).Delete(message).Error
}
func ShowAllMessageTemplete() []model.MessageTemplate {
	var messages []model.MessageTemplate
	DB.Model(&model.TemplateMessage{}).Find(&messages)
	return messages
}
