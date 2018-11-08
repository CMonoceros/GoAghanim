package http

import (
	"cmonoceros.com/GoAghanim/src/base"
	"cmonoceros.com/GoAghanim/src/http/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	router := gin.New()
	conf := base.GetDefaultConfig()

	routes.ApiV1PublicInit(router)
	routes.AdminPublicInit(router)
	routes.ApiV1Init(router)
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	router.Run(":" + conf.Port)
}
