package router

import (
	"cmonoceros.com/GoAghanim/pkg/controller"
	"cmonoceros.com/GoAghanim/pkg/lib"
	"cmonoceros.com/GoAghanim/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func ApiV1PublicInit(router *gin.Engine) {
	route := router.Group(
		"/api/v1",
		middleware.LogMiddleware(lib.API),
		middleware.HostMiddleware("api"))

	apiV1PublicInitGet(route)
	apiV1PublicInitPost(route)
}

func apiV1PublicInitPost(group *gin.RouterGroup) {
	group.POST("/login", controller.UserLogin)
}

func apiV1PublicInitGet(group *gin.RouterGroup) {
	group.GET("/checkup", func(c *gin.Context) {
		lib.Log.Info(`This is a checkup: check info log`)
		lib.Log.Error(`This is a checkup: check error log`)
		lib.GetContext(c).PackOK()
	})
}
