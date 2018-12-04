package router

import (
	"cmonoceros.com/GoAghanim/pkg/lib"
	"cmonoceros.com/GoAghanim/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func ApiV1Init(router *gin.Engine) {
	route := router.Group(
		"/api/v1",
		middleware.LogMiddleware(lib.API),
		middleware.HostMiddleware("api"),
		middleware.AuthMiddleware())

	apiV1InitGet(route)
	apiV1InitPost(route)
}

func apiV1InitPost(group *gin.RouterGroup) {
}

func apiV1InitGet(group *gin.RouterGroup) {
}
