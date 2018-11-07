package routes

import (
	"github.com/gin-gonic/gin"
	"jojotu.com/base"
	"jojotu.com/http/middlewares"
	"jojotu.com/services/user/controllers"
)

func ApiV1PublicInit(router *gin.Engine) {
	route := router.Group(
		"/api/v1",
		middlewares.LogMiddleware(base.API),
		middlewares.HostMiddleware("api"))

	apiV1PublicInitPost(route)
}

func apiV1PublicInitPost(group *gin.RouterGroup) {
	group.POST("/login", controllers.UserLogin)
}
