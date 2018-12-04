package middleware

import (
	"cmonoceros.com/GoAghanim/pkg/lib"
	"github.com/gin-gonic/gin"
)

func LogMiddleware(file lib.LogFile) gin.HandlerFunc {
	return func(c *gin.Context) {
		lib.InitLogger(file)
		lib.RequestLog.Info("-----------Start Request--------------------")
		lib.RequestLog.Info(c.Request)
		c.Next()
	}
}
