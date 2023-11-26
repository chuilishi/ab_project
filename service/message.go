package service

import (
	"ab_project/global"
	"ab_project/model"
	"ab_project/mysqlDB"
	"ab_project/service/message"
	"ab_project/service/response"
	"ab_project/wechat"
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"strconv"
)

// EditUserStatus 编辑用户状态 已经弃用
func EditUserStatus(c *gin.Context) {
	var userList = struct {
		id     []int  //id
		status string //下一状态
	}{}
	err := c.ShouldBind(&userList)
	if err != nil {
		response.FailWithMessage("参数请求错误"+err.Error(), c)
		return
	}
	var userErr []error
	for _, userId := range userList.id {
		err = mysqlDB.ChangeUserStatus(userId, userList.status)
		if err != nil {
			userErr = append(userErr, err)
		}
	}
	if len(userErr) == 0 {
		response.Ok(c)
		return
	}
	response.FailWithMessage("用户未能成功修改状态", c)
}

// OverComeProblems 解决用户异常信息
func OverComeProblems(c *gin.Context) {
	studentid := c.Query("studentid")
	if studentid == "" {
		response.FailWithMessage("参数请求错误", c)
		return
	}
	user := mysqlDB.IsUserHaveByStudentId(studentid)
	if user.ID == 0 {
		response.FailWithMessage("查无此人", c)
		return
	}
	err := mysqlDB.OverComeProblem(user)
	if err != nil {
		response.FailWithMessage("修改异常信息失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("修改异常信息成功", c)
}

// SendMessage 给用户发送信息
func SendMessage(c *gin.Context) {
	var msg = model.TemplateMessage{}
	err := c.ShouldBind(&msg)
	if err != nil {
		response.FailWithMessage("参数获取错误", c)
		return
	}
	re := regexp.MustCompile(`(\d{16})`)
	fmt.Println(msg.HTTP)
	if !re.MatchString(msg.HTTP) {
		response.FailWithMessage("无正确的HTTP参数", c)
		return
	}
	fmt.Println(msg)
	messageid := re.FindStringSubmatch(msg.HTTP)

	err = wechat.SendTemplateMessage(msg)
	if err != nil {
		response.FailWithMessage("发送微信消息失败"+err.Error(), c)
		return
	}

	err = message.SendMessage(msg.WxOpenId, msg.Msg, messageid[0])
	if err != nil {
		response.FailWithMessage("发送信息失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("消息发送获取成功", c)
	return
}

// IsIdUsed 给用户发信息的id是否已经使用
func IsIdUsed(c *gin.Context) {
	id := c.Query("id")
	fmt.Println(id)
	flag := mysqlDB.IsIdDUsed(id)
	if flag {
		response.FailWithMessage("该id已经使用", c)
		return
	}
	response.OkWithMessage("id未使用", c)
	return

}

// UserFeedBack code 1 通过，code2拒绝，code3异常信息
func UserFeedBack(c *gin.Context) {
	var msg struct {
		Code     int    `json:"code"`     //code == 1 同意 2 拒绝 3 异常
		Message  string `json:"message"`  //异常之后的信息
		Wxopenid string `json:"wxopenid"` //微信openid
		Status   string `json:"status"`   //code为1或2时更新用户状态
		ID       string `json:"id"`       //消息唯一id码
	}
	err := c.ShouldBind(&msg)
	if err != nil {
		response.FailWithMessage("无法获取到参数"+err.Error(), c)
		return
	}
	if msg.Code == 0 {
		response.FailWithMessage("参数请求错误", c)
		return
	}
	flag := mysqlDB.IsIdDUsed(msg.ID)
	if flag {
		response.FailWithMessage("该id已经使用", c)
		return

	}
	flag = mysqlDB.IsIDSend(msg.ID)
	if !flag {
		response.FailWithMessage("该消息id还未发出", c)
		return

	}
	user := mysqlDB.IsUserHave(msg.Wxopenid)
	if user.ID == 0 {
		response.FailWithMessage("无法查找到该用户", c)
		return
	} else if msg.Code == 1 && msg.Status != "不修改" && user.Status != msg.Status {
		err := message.PassMessage(msg.Wxopenid, msg.Status, msg.ID)
		if err != nil {
			response.FailWithMessage("记录失败无法更新数据库"+err.Error(), c)
			return
		}
		user.Status = msg.Status
		err = mysqlDB.UpdateUser(user)
		if err != nil {
			response.FailWithMessage("无法更新用户状态"+err.Error(), c)
			return
		}
		response.OkWithMessage("用户状态更新成功", c)
	} else if msg.Code == 1 && msg.Status == "不修改" {
		err := message.PassMessage(msg.Wxopenid, "不修改状态", msg.ID)
		if err != nil {
			response.FailWithMessage("记录失败无法更新数据库"+err.Error(), c)
			return
		}
		response.OkWithMessage("用户状态更新成功", c)

	} else if msg.Code == 2 && msg.Status != "不修改" && user.Status != msg.Status {
		err := message.NotPassMassage(msg.Wxopenid, msg.Status, msg.ID)
		if err != nil {
			response.FailWithMessage("记录失败无法更新数据库"+err.Error(), c)
			return
		}
		user.Status = msg.Status
		err = mysqlDB.UpdateUser(user)
		if err != nil {
			response.FailWithMessage("无法更新用户状态"+err.Error(), c)
			return
		}
		response.OkWithMessage("用户状态更新成功", c)
	} else if msg.Code == 2 && msg.Status == "不修改" {
		err := message.NotPassMassage(msg.Wxopenid, "不修改状态", msg.ID)
		if err != nil {
			response.FailWithMessage("记录失败无法更新数据库"+err.Error(), c)
			return
		}
		response.OkWithMessage("用户状态更新成功", c)
	} else if msg.Code == 3 {
		err := message.Problems(msg.Wxopenid, msg.Message, msg.ID)
		if err != nil {
			response.FailWithMessage("反馈信息失败"+err.Error(), c)
			return
		}
		response.OkWithMessage("反馈成功", c)
	}

	return
}

// UserPassHistory 用户通过历史信息
func UserPassHistory(c *gin.Context) {
	wxopenid := c.Query("wxopenid")
	if wxopenid == "" {
		response.FailWithMessage("参数请求错误", c)
		return
	}
	var find = []int{0, 3, 4}
	messages, err := mysqlDB.FindUserPassHistory(wxopenid, find)
	if err != nil {
		response.FailWithMessage("无法查找到记录", c)
		return
	}
	lengh := len(messages)
	user := mysqlDB.IsUserHave(wxopenid)
	if user.ID == 0 {
		response.FailWithMessage("无法查找到用户", c)
		return
	}
	if user.ISProblem == 1 {
		messages = append(messages, model.Message{
			Message: "异常状态待处理",
		})
		for i := global.Status2int[user.Status] + 1; i != 12; i++ {
			messages = append(messages, model.Message{
				Message: global.Allstatus[i],
			})

		}
	} else if user.Status == "筛选未通过" || user.Status == "初试未通过" || user.Status == "复试未通过" || user.Status == "终试未通过" {
	} else {
		for i := global.Status2int[user.Status] + 1; i != 12; i++ {
			messages = append(messages, model.Message{
				Message: global.Allstatus[i],
			})

		}
	}
	response.OkWithDetailed(messages, strconv.Itoa(lengh+user.ISProblem), c)
	return
}

// ShowUserHistory 展示用户的历史信息
func ShowUserHistory(c *gin.Context) {
	studentid := c.Query("studentid")
	student := mysqlDB.IsUserHaveByStudentId(studentid)
	if student.ID == 0 {
		response.FailWithMessage("无法查找到用户", c)
		return
	}
	messages, err := message.ShowMessage(studentid)
	if err != nil {
		response.FailWithMessage("无法查找到相关信息"+err.Error(), c)
		return
	}
	response.OkWithDetailed(messages, "查找成功", c)
	return
}

// SaveMessageTemplate 保存消息模版
func SaveMessageTemplate(c *gin.Context) {
	var messagetemplate model.MessageTemplate
	err := c.ShouldBind(&messagetemplate)
	if err != nil {
		response.FailWithMessage("参数请求错误"+err.Error(), c)
		return
	}
	err = mysqlDB.SaveMessageTemplate(&messagetemplate)
	if err != nil {
		response.FailWithMessage("无法保存消息"+err.Error(), c)
		return
	}
	response.OkWithMessage("保存消息模板成功", c)
	return
}

// DeleteMessageTemplate 删除消息模板
func DeleteMessageTemplate(c *gin.Context) {
	idd := c.Query("id")
	id, err := strconv.Atoi(idd)
	if err != nil {
		response.FailWithMessage("参数请求失败", c)
		return
	}
	err = mysqlDB.DeleteMessageTemplate(id)
	if err != nil {
		response.FailWithMessage("删除失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
	return
}

// ShowAllMessageTemplate 展示所有消息模板
func ShowAllMessageTemplate(c *gin.Context) {
	messages := mysqlDB.ShowAllMessageTemplete()
	response.OkWithDetailed(messages, "请求成功", c)
	return
}
