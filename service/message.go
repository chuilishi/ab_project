package service

import (
	"ab_project/mysqlDB"
	"ab_project/service/message"
	"ab_project/service/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// EditUserStatus 编辑用户状态
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

// IHaveProblems 用户反馈异常信息
func IHaveProblems(c *gin.Context) {
	var user = struct {
		Wxopenid string `json:"wxopenid"`
		Problem  string `json:"problem"`
	}{}
	err := c.ShouldBind(&user)
	fmt.Println(user)
	if err != nil {
		response.FailWithMessage("参数请求错误", c)
	}
	err = message.Problems(user.Wxopenid, user.Problem)
	if err != nil {
		response.FailWithMessage("反馈异常"+err.Error(), c)
		return
	}
	response.OkWithMessage("反馈成功！", c)
	return
}
func SendMessage(c *gin.Context) {
	var msg = struct {
		Wxopenid string `json:"wxopenid"`
		Message  string `json:"problem"`
	}{}
	err := c.ShouldBind(&msg)
	if err != nil {
		response.FailWithMessage("参数获取错误", c)
		return
	}
	err = message.AddMessage(msg.Wxopenid, msg.Wxopenid, 2)
	if err != nil {
		response.FailWithMessage("发送信息失败", c)
		return
	}

	//TODO HAHA

	response.OkWithMessage("消息发送获取成功", c)
	return
}
