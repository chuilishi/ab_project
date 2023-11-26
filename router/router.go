package router

import (
	_ "ab_project/docs"
	"ab_project/middle"
	"ab_project/service"
	"fmt"
	"github.com/gin-gonic/gin"
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
	r.Use(middle.Cors(), middle.Logger())
	r.POST("/postUserMessage", service.PostUserMessage)
	r.GET("/isUserExist", service.IsUserExist)
	r.POST("/uploadUserFileMessage", service.UploadUserFileMessage)
	r.DELETE("/deleteUserFileMessage", service.DeleteUserFileMessage)
	r.GET("/showUserFileMessage", service.ShowUserFileMessage)
	r.POST("/loginManager", service.LoginManage)
	r.POST("/feedback", service.UserFeedBack)
	r.GET("/allStatus", service.AllStatus)
	r.GET("/userHistory", service.UserPassHistory)
	r.GET("/homePicture", service.ShowHomePicture)
	r.GET("/isIdUsed", service.IsIdUsed)
	admin := r.Group("admin").Use(middle.JWTCheck())
	admin.POST("/updateUserMessage", service.AdminPostUserMessage)
	admin.GET("/showUserFileMessage", service.ShowUserFileMessage)
	admin.GET("/userDirection", service.FindUsersByDirection)
	admin.GET("/userStatus", service.FindUsersByStatus)
	admin.GET("/userProblems", service.FindProblemUsers)
	admin.GET("/showUserMessage", service.ShowUserMessage)
	admin.GET("/overComeProblem", service.OverComeProblems)
	admin.POST("/uploadPicture", service.UploadHomePicture)
	admin.POST("/sendMessageToUser", service.SendMessage)
	admin.GET("/showUserHistory", service.ShowUserHistory)
	admin.DELETE("/deleteMessageTemplate", service.DeleteMessageTemplate)
	admin.POST("/saveMessageTemplate", service.SaveMessageTemplate)
	admin.GET("/showAllMessageTemplate", service.ShowAllMessageTemplate)

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swagger.Handler))

	return r
}
func InitGetUserFileMessageHandler() {
	http.Handle("/", http.FileServer(http.Dir("./")))

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("文件服务器启动失败" + err.Error())
	}

}
