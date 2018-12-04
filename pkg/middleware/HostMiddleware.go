package middleware

import (
	"cmonoceros.com/GoAghanim/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func HostMiddleware(domain string) gin.HandlerFunc {
	return func(c *gin.Context) {
		conf := pkg.GetDefaultConfig()
		reg := regexp.MustCompile(`\Q` + domain + `.` + conf.AppUrl + `\E.*`)
		if !reg.MatchString(c.Request.Host) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.Next()
	}
}
