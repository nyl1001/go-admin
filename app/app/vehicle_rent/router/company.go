package router

import (
	"go-admin/app/app/vehicle_rent/apis"
	"go-admin/core/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerCompaniesRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckCompaniesRouter)
}

// registerCompaniesRouter 注册需要认证的路由
func registerCompaniesRouter(v1 *gin.RouterGroup) {
	api := apis.Companies{}
	r := v1.Group("/app/vehicle_rent/company").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/all", api.GetAll)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.GET("/export", api.Export)
		r.GET("/tree", api.Get2Tree)
	}
}

// registerNoCheckCompaniesRouter
func registerNoCheckCompaniesRouter(v1 *gin.RouterGroup) {}
