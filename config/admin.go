package config

import (
	"github.com/goravel/framework/foundation"
)

func init() {
	config := foundation.App.MakeConfig()
	config.Add("admin", map[string]any{
		// 应用名称
		"app_name": config.Env("ADMIN_APP_NAME", "Goravel"),
		// 应用 logo
		"logo": config.Env("ADMIN_LOGO", "/admin-assets/logo.png"),
		// 默认头像
		"default_avatar": config.Env("ADMIN_DEFAULT_AVATAR", "/admin-assets/default-avatar.png"),
		"route": map[string]any{
			"prefix": config.Env("ADMIN_ROUTE_PREFIX", "/admin-api"),
			"domain": config.Env("ADMIN_DOMAIN"),
		},
		"auth": map[string]any{
			// 是否启用认证
			"enabled": config.Env("ADMIN_AUTH_ENABLED", true),
			// 签权守卫
			"guard": "admin",
			// 登录时是否使用验证码
			"login_captcha": config.Env("ADMIN_LOGIN_CAPTCHA", true),
			//白名单: 请求方式 + 注册时的路由(支持正则)
			"exclude": []string{
				"post:/login",
				"get:/logout",
				"get:/_settings",
			},
		},
		"permission": map[string]any{
			// 是否启用权限
			"enabled": config.Env("ADMIN_PERMISSION_ENABLED", true),
			//白名单: 请求方式 + 注册时的路由(支持正则)
			"exclude": []string{},
		},
		"layout": map[string]any{
			// 浏览器标题, 功能名称使用 %title% 代替
			"title": config.Env("ADMIN_SITE_TITLE", "%title% | Goravel Admin"),
			// 顶部导航栏
			"header": map[string]any{
				// 是否显示 [刷新] 按钮
				"refresh": config.Env("ADMIN_HEADER_REFRESH", true),
				// 是否显示 [暗色模式] 按钮
				"dark": config.Env("ADMIN_HEADER_DARK", true),
				// 是否显示 [全屏] 按钮
				"full_screen": config.Env("ADMIN_HEADER_FULL_SCREEN", true),
				// 是否显示 [多语言] 按钮
				"locale_toggle": config.Env("ADMIN_HEADER_LOCALE_TOGGLE", false),
				// 是否显示 [主题配置] 按钮
				"theme_config": config.Env("ADMIN_HEADER_THEME_CONFIG", true),
			},
			// 多语言选项
			"locale_options": map[string]string{
				"zh_CN": "简体中文",
				"en":    "English",
			},
			// 底部信息
			"footer": config.Env("ADMIN_FOOTER"),
		},
		// 是否显示 [开发者工具]
		"show_development_tools": config.Env("ADMIN_SHOW_DEVELOPMENT_TOOLS", true),
	})
}
