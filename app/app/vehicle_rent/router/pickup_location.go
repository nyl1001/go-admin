package router

import (
	"go-admin/app/app/vehicle_rent/apis"
	"go-admin/core/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerPickupLocationRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckPickupLocationRouter)
}

// registerPickupLocationRouter 注册需要认证的路由
func registerPickupLocationRouter(v1 *gin.RouterGroup) {
	api := apis.PickupLocation{}
	r := v1.Group("/app/vehicle_rent/pickup_location").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.GET("/export", api.Export)
	}
}

// registerNoCheckPickupLocationRouter
func registerNoCheckPickupLocationRouter(v1 *gin.RouterGroup) {}
