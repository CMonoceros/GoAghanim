package router

import (
	"cmonoceros.com/GoAghanim/pkg/middleware"
	"cmonoceros.com/GoAghanim/pkg/web"
	"github.com/gin-gonic/gin"
)

// APIV1Init 接口v1私有路由初始化
func APIV1Init(router *gin.Engine) {
	route := router.Group(
		"/api/v1",
		middleware.LogMiddleware(web.LogAPIConfig),
		middleware.HostMiddleware("api"),
		middleware.AuthMiddleware())

	apiV1InitGet(route)
	apiV1InitPost(route)
}

func apiV1InitPost(group *gin.RouterGroup) {
}

func apiV1InitGet(group *gin.RouterGroup) {
}
