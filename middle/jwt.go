package middle

import (
	"ab_project/global"
	"ab_project/service"
	"ab_project/service/response"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// jwt中间件，用于检测当前登录用户是否合法
func JWTCheck() gin.HandlerFunc {
	// 我们这里jwt鉴权取头部信息 jwt-code 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
	return func(c *gin.Context) {
		y := c.Request.Header.Get("Jwt-Code")
		if y == "" {
			response.CannotPassWithMessage("用户未登录", c)
			c.Abort()
			return
		}
		token, err := jwt.ParseWithClaims(y, &service.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.JWTKey), nil
		})
		if token.Valid {
			//fmt.Println(token.Claims.(*service.MyCustomClaims).UserName)
			c.Next()
		} else {
			fmt.Println(err)
			response.CannotPassWithMessage("用户登录过期或非法登录", c)
			c.Abort()
			return
		}

	}
}
