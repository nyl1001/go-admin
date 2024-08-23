package constant

const (
	//app下载方式
	SysConfIsFrontend = "1"     //前台显示
	RoleKeyAdmin      = "admin" //角色类型，超级管理员

	//角色状态
	//RoleStatusStart = "1" //启用
	//RoleStatusStop  = "2" //停用

	//登录退出
	UserLoginStatus  = "1"
	UserLogoutStatus = "2"

	//数据权限类型
	DataScope1 = "1" //全部数据
	DataScope2 = "2" //自定数据权限
	DataScope3 = "3" //本部门数据权限
	DataScope4 = "4" //本部门及以下数据权限
	DataScope5 = "5" //仅本人数据权限

	// MenuM 目录
	MenuM string = "1"
	// MenuC 菜单
	MenuC string = "2"
	// MenuF 按钮
	MenuF string = "3"
)

// 模板相关
const (
	RouterName         = "router.go"
	BusinessRouterName = "businessRouter.go"
	ApiName            = "api.go"
	ModelName          = "model.go"
	DtoName            = "dto.go"
	ServiceName        = "service.go"
	JsName             = "js"
	VueName            = "index.vue"
	LangName           = "lang.go"
	ConstantName       = "constant.go"
	TemplateNameKey    = "name"
	TemplatePathKey    = "path"
)

// 模板相关
var (
	TemplatInfo = map[string]string{
		ModelName:          "static/template/model.go.template",
		DtoName:            "static/template/dto.go.template",
		ServiceName:        "static/template/service.go.template",
		ApiName:            "static/template/apis.go.template",
		BusinessRouterName: "static/template/business_router.go.template",
		JsName:             "static/template/js.go.template",
		VueName:            "static/template/vue.go.template",
		RouterName:         "static/template/router.go.template",
		LangName:           "static/template/lang.go.template",
		ConstantName:       "static/template/constant.go.template",
	}

	TemplateInfoList = []map[string]string{
		{
			TemplateNameKey: ModelName,
			TemplatePathKey: "static/template/model.go.template",
		},
		{
			TemplateNameKey: DtoName,
			TemplatePathKey: "static/template/dto.go.template",
		},
		{
			TemplateNameKey: ServiceName,
			TemplatePathKey: "static/template/service.go.template",
		},
		{
			TemplateNameKey: ApiName,
			TemplatePathKey: "static/template/apis.go.template",
		},
		{
			TemplateNameKey: BusinessRouterName,
			TemplatePathKey: "static/template/business_router.go.template",
		},
		{
			TemplateNameKey: JsName,
			TemplatePathKey: "static/template/js.go.template",
		},
		{
			TemplateNameKey: VueName,
			TemplatePathKey: "static/template/vue.go.template",
		},
		{
			TemplateNameKey: RouterName,
			TemplatePathKey: "static/template/router.go.template",
		},
		{
			TemplateNameKey: LangName,
			TemplatePathKey: "static/template/lang.go.template",
		},
		{
			TemplateNameKey: ConstantName,
			TemplatePathKey: "static/template/constant.go.template",
		},
	}
)
