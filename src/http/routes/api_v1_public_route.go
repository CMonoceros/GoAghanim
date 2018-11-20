package routes

import (
	"cmonoceros.com/GoAghanim/src/base"
	"cmonoceros.com/GoAghanim/src/http/controllers/api/v1"
	"cmonoceros.com/GoAghanim/src/http/middlewares"
	"cmonoceros.com/GoAghanim/src/http/response"
	"github.com/gin-gonic/gin"
)

func ApiV1PublicInit(router *gin.Engine) {
	route := router.Group(
		"/api/v1",
		middlewares.LogMiddleware(base.API),
		middlewares.HostMiddleware("api"))

	apiV1PublicInitGet(route)
	apiV1PublicInitPost(route)
}

func apiV1PublicInitPost(group *gin.RouterGroup) {
	group.POST("/login", v1.UserLogin)
}

func apiV1PublicInitGet(group *gin.RouterGroup) {
	group.GET("/checkup", func(c *gin.Context) {
		base.Log.Info(`This is a checkup: check info log`)
		base.Log.Error(`This is a checkup: check error log`)
		response.GetContext(c).PackOK()
	})
}
