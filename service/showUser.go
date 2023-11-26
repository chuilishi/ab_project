package service

import (
	"ab_project/mysqlDB"
	"ab_project/service/response"
	"github.com/gin-gonic/gin"
)

// ShowUserMessage 展示所有用户的简历投递信息
func ShowUserMessage(c *gin.Context) {
	users := mysqlDB.FindUsersByDirection("全部")
	countAll := 0
	allMessage := make(map[string][]interface{})
	sexMessage := make(map[string]int)
	gardMessage := make(map[string]int)
	directionMessage := make(map[string]int)
	timeMessage := make(map[string]int)
	professionMessage := make(map[string]int)
	statusMessage := make(map[string]int)
	for _, user := range users {
		countAll++
		sexMessage[user.Sex]++
		gardMessage[user.Grade]++
		directionMessage[user.Direction]++
		timeMessage[user.UpdatedAt.String()[0:11]]++
		professionMessage[user.Profession]++
		statusMessage[user.Status]++
	}
	countalll := make(map[string]int)
	countalll["count"] = countAll
	var count1 []int
	var name1 []string
	for key, value := range directionMessage {
		count1 = append(count1, value)
		name1 = append(name1, key)
	}
	allMessage["方向"] = append(allMessage["方向"], count1, name1)
	var count2 []int
	var name2 []string

	for key, value := range professionMessage {
		count2 = append(count2, value)
		name2 = append(name2, key)
	}
	allMessage["专业"] = append(allMessage["专业"], count2, name2)
	var count3 []int
	var name3 []string

	for key, value := range sexMessage {
		count3 = append(count3, value)
		name3 = append(name3, key)

	}
	allMessage["性别"] = append(allMessage["性别"], count3, name3)

	var count4 []int
	var name4 []string

	for key, value := range gardMessage {
		count4 = append(count4, value)
		name4 = append(name4, key)
	}

	allMessage["年级"] = append(allMessage["年级"], count4, name4)
	var count5 []int
	var name5 []string

	for key, value := range timeMessage {
		count5 = append(count5, value)
		name5 = append(name5, key)
	}
	allMessage["投递时间"] = append(allMessage["投递时间"], count5, name5)
	var count6 []int
	var name6 []string

	for key, value := range statusMessage {
		count6 = append(count6, value)
		name6 = append(name6, key)
	}

	allMessage["状态"] = append(allMessage["状态"], count6, name6)

	response.OkWithDetailed(allMessage, "用户信息", c)
}

// AllStatus 展示用户状态
func AllStatus(c *gin.Context) {

	return
}
