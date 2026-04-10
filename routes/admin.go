package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"github.com/khs1001/gowl-admin/http/controllers"
	"github.com/khs1001/gowl-admin/support/core"
)

func Admin(router route.Router) {
	facades.Route().Static("/admin-assets", "./public/admin-assets")
	core.IndexRoute(router)
	router.Prefix(core.ApiPreix()).Group(func(router route.Router) {
		// 系统相关
		index := controllers.NewIndexController()
		router.Get("_settings", index.Settings)      // 获取系统配置
		router.Post("_settings", index.SaveSettings) // 保存系统配置

		// 登录相关
		auth := controllers.NewAuthController()
		router.Post("login", auth.Login)             // 登录
		router.Get("logout", auth.Logout)            // 退出登录
		router.Get("current-user", auth.CurrentUser) // 获取当前用户信息
		router.Get("menus", auth.Menus)              // 获取菜单信息
	})
}
