package main

import (
	_ "ab_project/docs"
	"ab_project/wechat"
)

func main() {
	//mysqlDB.InitGrom()
	//r := router.GetRouter()
	//r.Run(":8080")
	wechat.Wechat()
}
