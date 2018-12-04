package middleware

import (
	"cmonoceros.com/GoAghanim/pkg/web/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.GetQuery("api_token")
		if err == false {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		isAccess, account := model.Auth(token)
		if !isAccess {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.Set("user", account)
			c.Next()
		}
	}
}
