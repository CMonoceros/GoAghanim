package routes

import (
	"github.com/gin-gonic/gin"
	"jojotu.com/base"
	"jojotu.com/http/middlewares"
)

func AdminPublicInit(router *gin.Engine) {
	route := router.Group(
		"",
		middlewares.LogMiddleware(base.ADMIN),
		middlewares.HostMiddleware("admin"),
	)

	adminPublicInitPost(route)
}

func adminPublicInitPost(group *gin.RouterGroup) {
	group.GET("/login", func(context *gin.Context) {
		base.Log.Info("bbbb")
	})
}
