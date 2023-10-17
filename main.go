package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	fmt.Println("nihao")

	r.Run("8080")
}
