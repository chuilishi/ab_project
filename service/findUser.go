package service

import (
	"ab_project/mysqlDB"
	"ab_project/service/response"
	"github.com/gin-gonic/gin"
)

// 通过方向查找学生
func FindUsersByDirection(c *gin.Context) {
	direction := c.Query("direction")
	users := mysqlDB.FindUsersByDirection(direction)

	response.OkWithDetailed(users, direction, c)

}
