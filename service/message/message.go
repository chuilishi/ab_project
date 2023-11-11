package message

import (
	"ab_project/global"
	"ab_project/mysqlDB"
	"ab_project/utils"
	"errors"
)

// AddMessage 	给用户发送消息，或者用户反馈消息
// code ： 1  //用户异常信息记录
// code ： 2  //给用户发过的信息记录
// code :  3  //用户反馈信息记录
// code :  0  //用户通过信息记录
func AddMessage(wxopenid, problem string, code int) error {
	user := mysqlDB.IsUserHave(wxopenid)
	if user.ID == 0 {
		return errors.New("无法查找到该学生")
	}
	err := mysqlDB.SendMessageToUser(wxopenid, problem, code)
	if err != nil {
		return err
	}
	return nil
}

// Problems 用户提交错误信息
func Problems(wxopenid, problem string) error {
	err := AddMessage(wxopenid, problem, 1)
	if err != nil {
		return err
	}
	err = mysqlDB.PostProblem(wxopenid, problem)
	if err != nil {
		return errors.New("反馈异常信息失败" + err.Error())
	}
	us := mysqlDB.IsUserHave(wxopenid)
	return utils.SendMail(global.MailTOManager, us.UserName)
}

// SendMessage todo
func SendMessage(wxopenid, problem string) error {
	err := AddMessage(wxopenid, problem, 1)
	if err != nil {
		return err
	}
	return nil
}
