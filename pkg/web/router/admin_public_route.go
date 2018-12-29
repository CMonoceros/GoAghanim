package router

import (
	"cmonoceros.com/GoAghanim/pkg/middleware"
	"cmonoceros.com/GoAghanim/pkg/web"
	"github.com/gin-gonic/gin"
)

// AdminPublicInit 后台公共路由初始化
func AdminPublicInit(router *gin.Engine) {
	route := router.Group(
		"",
		middleware.LogMiddleware(web.LogAdminConfig),
		middleware.HostMiddleware("admin"))

	adminPublicInitPost(route)
}

func adminPublicInitPost(group *gin.RouterGroup) {
}
