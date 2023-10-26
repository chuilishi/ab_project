package service

import (
	"ab_project/model"
	"ab_project/mysqlDB"
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

// Delivery 用户注册实现
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
func Delivery(c *gin.Context) {
	var user = new(model.User)
	var information = new(model.Information)
	c.ShouldBind(user)
	c.ShouldBind(information)

	fmt.Println(user)
	fmt.Println(information)
	_, err := GiveJWT()
	//fmt.Println(token, err)
	if err != nil {
		fmt.Println(err.Error())
	}
	return

}
