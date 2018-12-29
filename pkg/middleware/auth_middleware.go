package middleware

import (
	"cmonoceros.com/GoAghanim/pkg/web/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware 验证合法用户中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.GetQuery("api_token")
		if err == false {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		isAccess, account := model.UserAuth(token)
		if !isAccess {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.Set("user", account)
			c.Next()
		}
	}
}
