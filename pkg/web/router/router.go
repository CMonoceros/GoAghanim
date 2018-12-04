package router

import (
	"cmonoceros.com/GoAghanim/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	router := gin.New()
	conf := pkg.GetDefaultConfig()

	ApiV1PublicInit(router)
	AdminPublicInit(router)
	ApiV1Init(router)
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	router.Run(":" + conf.Port)
}
