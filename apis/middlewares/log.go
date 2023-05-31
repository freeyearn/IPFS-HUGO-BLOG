package middlewares

import (
	logs "IPFS-Blog-Hugo/utils/loggers"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type responseJsonWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseJsonWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var log = logs.GetLogger()
		w := &responseJsonWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 执行时间
		latencyTime := time.Since(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 返回格式
		returnJson := w.body.String()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		log.WithFields(logrus.Fields{
			"client_ip":    clientIP,
			"req_uri":      reqUri,
			"req_method":   reqMethod,
			"status_code":  statusCode,
			"return_json":  returnJson,
			"latency_time": latencyTime,
		}).Info()
	}
}
