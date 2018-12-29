package middleware

import (
	"cmonoceros.com/GoAghanim/pkg/lib"
	"github.com/gin-gonic/gin"
)

// LogMiddleware 配置日志中间件
func LogMiddleware(config lib.LogFileConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		lib.InitLogger(config)
		c.Next()
	}
}
