package router

import (
	"ab_project/service"
	"github.com/gin-gonic/gin"
	swagger "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// GetRouter 返回集成router
func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swagger.Handler))
	r.GET("/login", service.LoginUser)
	r.GET("/login/register", service.RegisterUser)
	return r
}
