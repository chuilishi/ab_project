package router

import (
	_ "ab_project/docs"
	"ab_project/middle"
	"ab_project/service"
	"github.com/gin-gonic/gin"
	swagger "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// GetRouter 返回集成router
//
//	@title			AB迎新系统接口文档
//	@version		1.0
//	@description	接口文档
//	@host			123.207.73.185:8080
func GetRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middle.Cors())
	r.GET("/isUserExist", service.IsUserExist)
	r.GET("/userDirection", service.FindUsersByDirection)
	r.POST("/postUserMessage", service.PostUserMessage)
	r.POST("/createUserMessage", service.PostUserMessage)
	r.POST("sendMessageToUser", service.SendMessageToUser)
	r.POST("/upload", service.Upload)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swagger.Handler))
	return r
}
