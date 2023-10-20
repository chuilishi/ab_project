package service

import (
	_ "ab_project/docs"
	"ab_project/mysqlDB"
	"fmt"
	"github.com/gin-gonic/gin"
)

// LoginUser 	登录用户实现
// @Summary 	用户登录请求
// @Description 用户登录请求
// @Tags 		用户服务
// @Accept 		json
// @Produce 	json
// @Param 		username 	query 	string 	true "用户id"
// @Param 		password 	query 	string 	true "用户密码"
// @Success     200   	 	{object}    	people.User
// @Failure 	400  		{string}json   	"{"code":0,"msg": "查无此人"}"
// @Router 		/login [get]
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
// @Summary 	用户注册请求
// @Description 用户注册请求
// @Tags 		用户服务
// @Accept 		json
// @Produce 	json
// @Param 		username 	query 	string 	true "用户id"
// @Param 		password 	query 	string 	true "用户密码"
// @Success     200   	 	string  	"{"code":0,"msg": "注册成功"}"
// @Failure 	400  		string  	"{"code":0,"msg": "注册失败"}"
// @Router 		/login/register [get]
func RegisterUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	err := mysqlDB.RegisterUserByUsername(username, password)
	fmt.Println(err)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    0,
			"message": "注册失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "注册成功",
	})
	return

}
