package service

import (
	"ab_project/mysqlDB"
	"ab_project/service/response"
	"github.com/gin-gonic/gin"
)

func ShowUserMessage(c *gin.Context) {
	users := mysqlDB.FindUsersByDirection("全部")
	countAll := 0
	allMessage := make(map[string](map[string]int))
	sexMessage := make(map[string]int)
	gardMessage := make(map[string]int)
	directionMessage := make(map[string]int)
	timeMessage := make(map[string]int)
	professionMessage := make(map[string]int)
	for _, user := range users {
		countAll++
		sexMessage[user.Sex]++
		gardMessage[user.Grade]++
		directionMessage[user.Direction]++
		timeMessage[user.UpdatedAt.String()[0:11]]++
		professionMessage[user.Profession]++
	}
	countalll := make(map[string]int)
	countalll["count"] = countAll
	allMessage["方向"] = directionMessage
	allMessage["专业"] = professionMessage
	allMessage["性别"] = sexMessage
	allMessage["年级"] = gardMessage
	allMessage["投递时间"] = timeMessage
	allMessage["总人数"] = countalll

	response.OkWithDetailed(allMessage, "用户信息", c)
}
