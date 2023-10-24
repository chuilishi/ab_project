package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Access-Control-Allow-Credentials=true和Access-Control-Allow-Origin="*"有冲突
		//故Access-Control-Allow-Origin需要指定具体得跨域origin或者直接Access-Control-Allow-Origin="*"
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "content-type")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
		//c.Header("Access-Control-Expose-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "")
			c.Abort()
			return
		}
		c.Next()
	}
}
