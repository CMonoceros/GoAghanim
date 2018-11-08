package middlewares

import (
	"cmonoceros.com/GoAghanim/src/base"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func HostMiddleware(domain string) gin.HandlerFunc {
	return func(c *gin.Context) {
		conf := base.GetDefaultConfig()
		reg := regexp.MustCompile(`\Q` + domain + `.` + conf.AppUrl + `\E.*`)
		if !reg.MatchString(c.Request.Host) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.Next()
	}
}
