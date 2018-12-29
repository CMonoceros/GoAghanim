package middleware

import (
	"cmonoceros.com/GoAghanim/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

// HostMiddleware 验证host中间件
func HostMiddleware(domain string) gin.HandlerFunc {
	return func(c *gin.Context) {
		reg := regexp.MustCompile(`\Q` + domain + `.` + web.ENV.Get("APP_URL") + `\E.*`)
		if !reg.MatchString(c.Request.Host) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.Next()
	}
}
