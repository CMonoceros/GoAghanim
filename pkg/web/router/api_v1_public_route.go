package router

import (
	"cmonoceros.com/GoAghanim/pkg/lib"
	"cmonoceros.com/GoAghanim/pkg/middleware"
	"cmonoceros.com/GoAghanim/pkg/web"
	"cmonoceros.com/GoAghanim/pkg/web/controller"
	"github.com/gin-gonic/gin"
)

// APIV1PublicInit 接口v1公共路由初始化
func APIV1PublicInit(router *gin.Engine) {
	route := router.Group(
		"/api/v1",
		middleware.LogMiddleware(web.LogAPIConfig),
		middleware.HostMiddleware("api"))

	apiV1PublicInitGet(route)
	apiV1PublicInitPost(route)
}

func apiV1PublicInitPost(group *gin.RouterGroup) {
	group.POST("/login", controller.UserLogin)
}

func apiV1PublicInitGet(group *gin.RouterGroup) {
	group.GET("/checkup", func(c *gin.Context) {
		lib.Logger.Info(`This is a checkup: check info log`)
		lib.Logger.Error(`This is a checkup: check error log`)
		lib.GetContext(c).PackOK()
	})
}
