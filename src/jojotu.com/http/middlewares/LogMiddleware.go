package middlewares

import (
	"github.com/gin-gonic/gin"
	"jojotu.com/base"
)

func LogMiddleware(file base.LogFile) gin.HandlerFunc {
	return func(c *gin.Context) {
		base.InitLogger(file)
		base.RequestLog.Info("-----------Start Request--------------------")
		base.RequestLog.Info(c.Request)
		c.Next()
	}
}
