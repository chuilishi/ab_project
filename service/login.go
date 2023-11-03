package service

import (
	"ab_project/global"
	"ab_project/model"
	"ab_project/mysqlDB"
	"ab_project/service/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// IsUserExist 查询用户是否已经提交过
// 未提交过返回0
// 已经提交过返回code 1并且返回用户结构体
func IsUserExist(c *gin.Context) {
	wxopenid := c.Query("openid")

	fmt.Println(wxopenid)
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
	fmt.Println(user)
	err = mysqlDB.RegisterUser(user)
	if err != nil {
		fmt.Println(err)
		response.FailWithMessage("简历投递失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("简历投递成功", c)
	return
}
func LoginManage(c *gin.Context) {
	manageID := c.Query("managername")
	password := c.Query("password")
	if manageID == global.ManageID && password == global.ManagePassword {
		response.OkWithMessage(global.ManageName, c)
		return
	}
	response.FailWithMessage("账号或密码错误", c)
}
