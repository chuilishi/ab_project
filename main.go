package main

import (
	_ "ab_project/docs"
	"ab_project/mysqlDB"
	"ab_project/router"
)

func main() {
	mysqlDB.InitGrom()
	r := router.GetRouter()
	r.Run(":8080")
}
