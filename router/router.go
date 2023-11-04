package router

import (
	_ "ab_project/docs"
	"ab_project/middle"
	"ab_project/service"
	"fmt"
	"github.com/gin-gonic/gin"
	swagger "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// GetRouter 返回集成router
//
//	@title			AB迎新系统接口文档
//	@version		1.0
//	@description	接口文档
//	@host			123.207.73.185:8080
func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/loginManager", service.LoginManage)
	r.Use(middle.Cors())

	r.GET("/isUserExist", service.IsUserExist)
	r.GET("/userDirection", service.FindUsersByDirection)
	r.POST("/postUserMessage", service.PostUserMessage)
	r.POST("/createUserMessage", service.PostUserMessage)
	r.POST("/sendMessageToUser", service.SendMessageToUser)
	r.POST("/uploadUserFileMessage", service.UploadUserFileMessage)
	r.DELETE("/deleteUserFileMessage", service.DeleteUserFileMessage)
	r.GET("/showUserFileMessage", service.ShowUserFileMessage)
	r.GET("/showUserMessage", service.ShowUserMessage)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swagger.Handler))

	return r
}
func InitGetUserFileMessageHandler() {
	http.Handle("/", http.FileServer(http.Dir("./")))

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("文件服务器启动失败" + err.Error())
	}

}
