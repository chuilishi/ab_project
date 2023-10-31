package service

import (
	"ab_project/mysqlDB"
	"ab_project/service/response"
	"github.com/gin-gonic/gin"
)

// FindUsersByDirection 通过方向查找学生
func FindUsersByDirection(c *gin.Context) {
	direction := c.Query("direction")
	users := mysqlDB.FindUsersByDirection(direction)

	response.OkWithDetailed(users, direction, c)

}

// FindUsersByStatus 通过状态查找学生
func FindUsersByStatus(c *gin.Context) {
	status := c.Query("status")
	users := mysqlDB.FindUsersByStatus(status)

	response.OkWithDetailed(users, status, c)

}
