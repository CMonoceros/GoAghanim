package router

import (
	"net/http"

	"cmonoceros.com/GoAghanim/pkg/lib"
	"cmonoceros.com/GoAghanim/pkg/web"
	"github.com/gin-gonic/gin"
)

// Init 路由初始化
func Init() {
	router := gin.New()

	router.Use(func(c *gin.Context) {
		lib.InitLogger(web.LogRequestConfig)
		lib.Logger.Info("-----------Start Request--------------------")
		lib.Logger.Info(c.Request)
		c.Next()
	})

	AdminPublicInit(router)
	APIV1PublicInit(router)
	APIV1Init(router)
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	err := router.Run(":" + web.ENV.Get("PORT"))
	lib.CheckAndThrowError(err, lib.DefaultCode)
}
