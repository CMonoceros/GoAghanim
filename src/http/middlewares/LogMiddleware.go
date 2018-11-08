package middlewares

import (
	"cmonoceros.com/GoAghanim/src/base"
	"github.com/gin-gonic/gin"
)

func LogMiddleware(file base.LogFile) gin.HandlerFunc {
	return func(c *gin.Context) {
		base.InitLogger(file)
		base.RequestLog.Info("-----------Start Request--------------------")
		base.RequestLog.Info(c.Request)
		c.Next()
	}
}
