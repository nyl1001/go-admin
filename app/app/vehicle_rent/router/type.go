package router

import (
	"go-admin/app/app/vehicle_rent/apis"
	"go-admin/core/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerVehicleTypeRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckVehicleTypeRouter)
}

// registerVehicleTypeRouter 注册需要认证的路由
func registerVehicleTypeRouter(v1 *gin.RouterGroup) {
	api := apis.VehicleType{}
	r := v1.Group("/app/vehicle_rent/type").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.GET("/export", api.Export)
	}
}

// registerNoCheckVehicleTypeRouter
func registerNoCheckVehicleTypeRouter(v1 *gin.RouterGroup) {}
