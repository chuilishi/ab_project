package middle

import (
	"ab_project/service/response"
	"github.com/gin-gonic/gin"
)

// jwt中间件，用于检测当前登录用户是否合法
func JWTCheck() gin.HandlerFunc {
	// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
	return func(c *gin.Context) {
		y := c.Request.Header.Get("jwt-code")
		if y == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或者非法访问", c)
			c.Abort()
		}

		c.Next()
	}
}
