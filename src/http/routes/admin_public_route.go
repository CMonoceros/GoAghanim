package routes

import (
	"cmonoceros.com/GoAghanim/src/base"
	"cmonoceros.com/GoAghanim/src/http/middlewares"
	"github.com/gin-gonic/gin"
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
}
