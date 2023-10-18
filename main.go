package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

func main() {
	r := gin.Default()
	//r.Use(gin.LoggerWithFormatter())
	r.Use(MiddleWare())
	r.GET("/")
	err := r.Run(":80")
	if err != nil {
		fmt.Println("error")
		return
	}
}

//
//func logger(params gin.LogFormatterParams) string {
//	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
//		params.ClientIP,
//		params.TimeStamp.Format(time.RFC3339Nano),
//		params.Method,
//		params.Path,
//		params.Request.Proto,
//		params.StatusCode,
//		params.Latency,
//		params.Request.UserAgent(),
//	)
//}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("asdfadsf")
		signature := c.Query("signature")
		timestamp := c.Query("timestamp")
		nonce := c.Query("nonce")
		token := "abproject"
		echostr := c.Query("echostr")
		fmt.Printf("signature=%s timestamp=%s nonce=%s echostr=%s", signature, timestamp, nonce, echostr)
		tmpArr := []string{token, timestamp, nonce}
		sort.Strings(tmpArr)
		tmpStr := sha1.Sum([]byte(fmt.Sprintf("%s%s%s", tmpArr[0], tmpArr[1], tmpArr[2])))
		calculatedSignature := fmt.Sprintf("%x", tmpStr)
		if calculatedSignature == signature {
			c.String(http.StatusOK, echostr)
			c.Next()
		} else {
			fmt.Println("无效")
		}
	}
}

//func main() {
//	go hello()
//	time.Sleep(1*time.)
//	fmt.Println("main function")
//}
