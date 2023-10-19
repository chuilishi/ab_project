package router

import (
	"ab_project/service"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/login", service.LoginUser)
	return r
}
