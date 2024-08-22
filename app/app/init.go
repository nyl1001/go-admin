package app

/**
 * app应用体量过大的时候，可以在该应用根目录下新建这样一个文件，将子模块的路由和多语言都辉总在这里，然后再去项目根目录(go-admin/app/init.go)中注册整个应用
 */
import (
	appLang "go-admin/app/app/user/lang"
	appUserRouter "go-admin/app/app/user/router"
	appCompanyRouter "go-admin/app/app/vehicle_rent/router"
)

// AllRouter
// @Description: 汇总各大板块接口
// @return []func()
func AllRouter() []func() {
	//初始化路由
	var routers []func()

	//app-用户管理
	routers = append(routers, appUserRouter.InitRouter, appCompanyRouter.InitRouter)
	return routers
}

// LangInit
// @Description: 多语言初始化
func LangInit() {
	appLang.Init()
}
