package router

import (
	"go-admin/app/app/vehicle_rent/apis"
	"go-admin/core/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerFleetRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckFleetRouter)
}

// registerFleetRouter 注册需要认证的路由
func registerFleetRouter(v1 *gin.RouterGroup) {
	api := apis.Fleet{}
	r := v1.Group("/app/vehicle_rent/fleet").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.GET("/export", api.Export)
	}
}

// registerNoCheckFleetRouter
func registerNoCheckFleetRouter(v1 *gin.RouterGroup) {}
