package routes

import (
	"cmonoceros.com/GoAghanim/src/base"
	"cmonoceros.com/GoAghanim/src/http/controllers/api/v1"
	"cmonoceros.com/GoAghanim/src/http/middlewares"
	"github.com/gin-gonic/gin"
)

func ApiV1PublicInit(router *gin.Engine) {
	route := router.Group(
		"/api/v1",
		middlewares.LogMiddleware(base.API),
		middlewares.HostMiddleware("api"))

	apiV1PublicInitPost(route)
}

func apiV1PublicInitPost(group *gin.RouterGroup) {
	group.POST("/login", v1.UserLogin)
}
