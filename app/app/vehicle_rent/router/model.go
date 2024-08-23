package router

import (
	"go-admin/app/app/vehicle_rent/apis"
	"go-admin/core/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerModelRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckModelRouter)
}

// registerModelRouter 注册需要认证的路由
func registerModelRouter(v1 *gin.RouterGroup) {
	api := apis.Model{}
	r := v1.Group("/app/vehicle_rent/model").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.GET("/export", api.Export)
	}
}

// registerNoCheckModelRouter
func registerNoCheckModelRouter(v1 *gin.RouterGroup) {}
