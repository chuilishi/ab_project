package main

import (
	_ "ab_project/docs"
	"ab_project/mysqlDB"
	"ab_project/router"
)

// @title AB迎新系统接口文档
// @version 1.0
// @description 接口文档
// @host 123.207.73.185:8080
func main() {
	mysqlDB.InitGrom()
	r := router.GetRouter()
	r.Run(":8080")
}
