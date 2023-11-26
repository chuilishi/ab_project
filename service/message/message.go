package message

import (
	"ab_project/global"
	"ab_project/model"
	"ab_project/mysqlDB"
	"ab_project/utils"
	"errors"
)

// AddMessage 	给用户发送消息，或者用户反馈消息
// code ： 1  //用户异常信息记录
// code ： 2  //给用户发过的信息记录
// code :  3  //用户不通过信息记录
// code :  0  //用户通过信息记录
// code :  4  //管理员手动修改用户状态记录
func AddMessage(wxopenid, msg string, code int, messageid string) error {
	user := mysqlDB.IsUserHave(wxopenid)
	if user.ID == 0 {
		return errors.New("无法查找到该学生")
	}
	err := mysqlDB.SendMessageToUser(wxopenid, msg, code, messageid)
	if err != nil {
		return err
	}
	return nil
}

// Problems 用户提交错误信息
func Problems(wxopenid, problem string, messageid string) error {
	err := AddMessage(wxopenid, problem, 1, messageid)
	if err != nil {
		return err
	}
	err = mysqlDB.PostProblem(wxopenid, problem)
	if err != nil {
		return errors.New("反馈异常信息失败" + err.Error())
	}
	us := mysqlDB.IsUserHave(wxopenid)
	return utils.SendMail(global.MailTOManager, us.UserName, problem)
}

// SendMessage 给用户发消息
func SendMessage(wxopenid, msg, messageid string) error {
	err := AddMessage(wxopenid, msg, 2, messageid)

	if err != nil {
		return err
	}
	return nil
}

// PassMessage 用户通过
func PassMessage(wxopenid, msg, messageid string) error {
	err := AddMessage(wxopenid, msg, 0, messageid)
	return err
}

// NotPassMassage 用户不通过
func NotPassMassage(wxopenid, msg, message string) error {
	err := AddMessage(wxopenid, msg, 3, message)
	return err
}

// ShowMessage 显示用户所有消息
func ShowMessage(studentid string) ([][]model.Message, error) {
	messages, err := mysqlDB.FindUserPassHistoryByStudentid(studentid, -1)
	var messagelist [][]model.Message
	for _, message := range messages {
		if message.Code == 4 || message.Code == 2 {
			var item []model.Message
			item = append(item, message)
			messagelist = append(messagelist, item)
		} else {
			for i, _ := range messagelist {
				if messagelist[i][0].MessageId == message.MessageId {
					messagelist[i] = append(messagelist[i], message)
					break
				}
			}
		}
	}
	return messagelist, err
}
