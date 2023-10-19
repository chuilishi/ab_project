package service

import (
	"ab_project/mysqlDB"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user, err := mysqlDB.FindUserByUsernamePassword(username, password)
	fmt.Println(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "0",
			"message": "查无此人",
		})
		return
	}
	//use, err := json.Marshal(user)
	//if err != nil {
	//	fmt.Println("序列化失败")
	//}
	c.JSON(200, user)
}
