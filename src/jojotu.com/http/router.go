package http

import (
	"github.com/gin-gonic/gin"
	"jojotu.com/base"
	"jojotu.com/http/routes"
	"net/http"
)

func Init() {
	router := gin.New()
	conf := base.GetDefaultConfig()

	routes.ApiV1PublicInit(router)
	routes.AdminPublicInit(router)
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	router.Run(":" + conf.Port)
}
