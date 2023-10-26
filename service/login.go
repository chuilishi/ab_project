package service

import (
	"ab_project/model"
	"ab_project/mysqlDB"
	"ab_project/service/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// LoginUser 	登录用户实现
//
//	@Summary		用户登录请求
//	@Description	用户登录请求
//	@Tags			用户服务
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string		true	"用户id"
//	@Param			password	query		string		true	"用户密码"
//	@Success		200			{object}	model.User	"用户结构体json"
//	@Failure		400			{object}	Response 	"错误信息"
//	@Router			/login [get]
func LoginUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user, err := mysqlDB.FindUserByUsernamePassword(username, password)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 0,
			"msg":  "查无此人",
		})
		return
	}

	c.JSON(200, Marshal(user))
}

// RegisterUser 用户注册实现
//
//	@Summary		用户注册请求
//	@Description	用户注册请求
//	@Tags			用户服务
//	@Accept			json
//	@Produce		json
//	@Param			username	query	string	true	"用户id"
//	@Param			password	query	string	true	"用户密码"
//	@Success		200			{object}	Response	"正确信息"
//	@Failure		400			{object}	Response	"错误信息"
//	@Router			/register [get]
func RegisterUser(c *gin.Context) {
	//var user model.User
	//c.BindJSON(&user)
	//err := mysqlDB.RegisterUserByUsername(&user)
	//fmt.Println(err)
	//if err != nil {
	//	c.JSON(400, gin.H{
	//		"code":    0,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//c.JSON(200, gin.H{
	//	"code":    1,
	//	"message": "注册成功",
	//})
	token, err := GiveJWT()
	fmt.Println(token, err)

	return

}
func Exist(c *gin.Context) {

}

// Register 用户注册实现
func Register(c *gin.Context) {
	user := &model.User{}
	err := c.ShouldBind(user)
	if err != nil {
		fmt.Println(err)
		response.FailWithMessage("简历投递失败", c)
		return
	}
	fmt.Println(user)
	token, err := GiveJWT()
	if err != nil {
		fmt.Println("jwt生成失败" + err.Error())
		return
	}

	us, err := mysqlDB.RegisterUser(user)
	if err != nil {
		fmt.Println(err)
		response.FailWithMessage("用户注册失败"+err.Error(), c)
		return
	}
	c.String(200, token)
	return
}
