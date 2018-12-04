package router

import (
	"cmonoceros.com/GoAghanim/pkg/lib"
	"cmonoceros.com/GoAghanim/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func AdminPublicInit(router *gin.Engine) {
	route := router.Group(
		"",
		middleware.LogMiddleware(lib.ADMIN),
		middleware.HostMiddleware("admin"),
	)

	adminPublicInitPost(route)
}

func adminPublicInitPost(group *gin.RouterGroup) {
}
