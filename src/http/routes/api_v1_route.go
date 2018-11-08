package routes

import (
	"cmonoceros.com/GoAghanim/src/base"
	"cmonoceros.com/GoAghanim/src/http/middlewares"
	"github.com/gin-gonic/gin"
)

func ApiV1Init(router *gin.Engine) {
	route := router.Group(
		"/api/v1",
		middlewares.LogMiddleware(base.API),
		middlewares.HostMiddleware("api"),
		middlewares.AuthMiddleware())

	apiV1InitGet(route)
	apiV1InitPost(route)
}

func apiV1InitPost(group *gin.RouterGroup) {
}

func apiV1InitGet(group *gin.RouterGroup) {
}
