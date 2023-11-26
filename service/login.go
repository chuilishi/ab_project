package service

import (
	"ab_project/global"
	"ab_project/model"
	"ab_project/mysqlDB"
	"ab_project/service/message"
	"ab_project/service/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// IsUserExist 查询用户是否已经提交过
// 未提交过返回0
// 已经提交过返回code 1并且返回用户结构体
func IsUserExist(c *gin.Context) {
	wxopenid := c.Query("wxopenid")
	if wxopenid == "" {
		response.FailWithMessage("参数请求错误", c)
	}
	user := mysqlDB.IsUserHave(wxopenid)
	if user.ID == 0 {
		response.FailWithMessage("用户没有提交过", c)
		return
	}
	response.OkWithDetailed(user, "用户已经提交过", c)
}

// PostUserMessage 用户简历投递实现
// 用户不存在那么创建，用户存在就更新
func PostUserMessage(c *gin.Context) {
	user := &model.User{}
	err := c.ShouldBind(user)
	if err != nil {
		fmt.Println(err)
		response.FailWithMessage("简历投递失败,"+err.Error(), c)
		return
	}
	if user.WxOpenId == "" {
		response.FailWithMessage("无法获得微信openid", c)
		return
	}
	if user.StudentId == "" {
		response.FailWithMessage("无法获取学号", c)
		return
	}
	if user.OK == 1 {
		user.Status = "提交成功"
	} else {
		user.Status = "草稿"
	}
	err = mysqlDB.RegisterUser(user)
	if err != nil {
		fmt.Println(err)
		response.FailWithMessage("简历投递失败"+err.Error(), c)
		return
	}
	var find = []int{0}
	messages, err := mysqlDB.FindUserPassHistory(user.WxOpenId, find)
	if err != nil {
		response.FailWithMessage("数据库出错"+err.Error(), c)
		return
	}
	if len(messages) == 0 {
		err = message.AddMessage(user.WxOpenId, "草稿", 4, "0")
		if err != nil {
			response.FailWithMessage("数据库出错"+err.Error(), c)
			return
		}
	}
	if user.OK == 1 {
		var find = []int{0}
		messages, err := mysqlDB.FindUserPassHistory(user.WxOpenId, find)
		if err != nil {
			response.FailWithMessage("数据库出错"+err.Error(), c)
			return
		}
		flag := 0
		for _, message := range messages {
			if message.Message == "待筛选" {
				flag = 1
			}
		}
		if flag == 0 {
			message.AddMessage(user.WxOpenId, "提交成功", 4, "0")
			message.AddMessage(user.WxOpenId, "待筛选", 4, "0")
			user.Status = "待筛选"

		}

		err = mysqlDB.UpdateUser(user)
		if err != nil {
			response.FailWithMessage("更新状态失败"+err.Error(), c)
		}

	}
	response.OkWithMessage("简历投递成功", c)
	return
}

// AdminPostUserMessage 管理员修改用户简历
func AdminPostUserMessage(c *gin.Context) {
	user := &model.User{}
	err := c.ShouldBind(user)
	if err != nil {
		fmt.Println(err)
		response.FailWithMessage("参数请求错误,"+err.Error(), c)
		return
	}
	tempuser := mysqlDB.IsUserHaveByStudentId(user.StudentId)
	if user.Status != "" && tempuser.Status != user.Status {
		message.AddMessage(tempuser.WxOpenId, user.Status, 4, "0")
	}
	err = mysqlDB.RegisterUser(user)
	if err != nil {
		fmt.Println(err)
		response.FailWithMessage("简历修改失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("简历修改成功", c)
	return
}

// LoginManage 管理员登录实现
func LoginManage(c *gin.Context) {
	var manager model.Manager
	c.ShouldBind(&manager)
	fmt.Println(manager.ManagerName)
	fmt.Println(manager.Password)
	if manager.ManagerName == global.ManageID && manager.Password == global.ManagePassword {
		jwttoken, err := GiveJWT(manager.ManagerName)
		if err != nil {
			fmt.Println("无法生成jwt-token", err)
			return
		}
		response.OkWithDetailed(gin.H{"jwtCode": jwttoken}, global.ManageName, c)
		return
	}
	response.FailWithMessage("账号或密码错误", c)
}
