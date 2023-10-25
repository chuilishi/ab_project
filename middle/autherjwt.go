package middle

import (
	"ab_project/global"
	"ab_project/service/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// jwt中间件，用于检测当前登录用户是否合法
func JWTCheck() gin.HandlerFunc {
	// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
	return func(c *gin.Context) {
		y := c.Request.Header.Get("jwt-code")
		if y == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "用户未登录", c)
			c.Abort()
		}
		_, err := jwt.Parse(y, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.JWTKey), nil
		})
		if err != nil {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{"reload": true}, "用户登录过期", c)
		}

	}
}
