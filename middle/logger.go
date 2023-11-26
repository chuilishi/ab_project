package middle

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// CustomResponseWriter 封装 gin ResponseWriter 用于获取回包内容。
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 打印请求信息
		start := time.Now()

		crw := &CustomResponseWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = crw
		reqBody, _ := c.GetRawData()
		zap.L().Info("request received ", zap.String("request ip", c.ClientIP()), zap.String("request Method", c.Request.Method), zap.String("requestURI", c.Request.RequestURI), zap.String("request Body", string(reqBody)))
		// 执行请求处理程序和其他中间件函数
		c.Next()

		// 记录回包内容和处理时间
		end := time.Now()
		latency := end.Sub(start)
		respBody := string(crw.body.Bytes())
		zap.L().Info("response send ", zap.String("request ip", c.ClientIP()), zap.String("last times", fmt.Sprintf("%v", latency)), zap.String("request Method", c.Request.Method), zap.String("requestURI", c.Request.RequestURI), zap.String("response Body", string(respBody)))

	}
}
