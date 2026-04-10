package core

import (
	"os"
	"regexp"
	"strings"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/str"
	"github.com/khs1001/gowl-admin/consts"
)

// IndexRoute 注册管理后台首页路由
// 参数 router：路由注册器实例，用于注册新的路由
func IndexRoute(router route.Router) {
	// 从配置中获取管理后台API前缀，如果未配置则使用默认值
	adminApiPrefix := facades.Config().GetString(consts.AdminRoutePrefix, consts.DefaultApiPrefix)
	// 生成管理后台首页的URL前缀（从API前缀中移除"-api"后缀）
	adminIndexPrefix := str.Of(adminApiPrefix).RTrim("-api").String()
	// 注册GET路由，用于提供管理后台首页
	facades.Route().Get(adminIndexPrefix, func(ctx http.Context) http.Response {
		// 读取管理后台首页HTML文件内容
		content, err := os.ReadFile(facades.App().PublicPath("admin-assets/index.html"))
		// 如果文件读取失败，返回404错误
		if err != nil {
			ctx.Request().Abort(http.StatusNotFound)
		}
		// 将HTML内容中的默认API前缀路径("/admin-api")替换为实际配置的API前缀
		content = []byte(str.Of(string(content)).Replace(consts.DefaultApiPrefix, adminApiPrefix).String())
		// 返回修改后的HTML内容作为响应，状态码为200
		return ctx.Response().Data(http.StatusOK, "text/html;", content)
	})
}

// ApiPreix 构建管理后台的API路径前缀
// 参数 parts：可选的路径部分，将被添加到基础API前缀之后
// 返回值：完整的API路径前缀字符串
func ApiPreix(parts ...string) string {
	adminApiPrefix := facades.Config().GetString(consts.AdminRoutePrefix, consts.DefaultApiPrefix)
	adminApiPrefix = strings.Join(append([]string{adminApiPrefix}, parts...), "/")
	return adminApiPrefix
}

// IsAdminRequest 判断请求是否为后台请求
func IsAdminRequest(originalURL string) bool {
	adminApiPrefix := facades.Config().GetString(consts.AdminRoutePrefix, consts.DefaultApiPrefix)
	return str.Of(originalURL).StartsWith(adminApiPrefix)
}

// IsAllowRequest 判断是否允许请求
func IsAllowRequest(rule, method, originalURL string) bool {
	matchedMethod := true
	cleanRule := strings.TrimLeft(rule, "^")
	methodList := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "CONNECT", "TRACE"}

	// 移除请求路径中的参数以及前缀
	url := strings.TrimLeft(strings.Split(originalURL, "?")[0], "/")
	url = strings.TrimLeft(url, strings.Split(url, "/")[0])
	url = strings.TrimLeft(url, "/")

	// 校验请求方式
	for _, v := range methodList {
		prefixUpper := v + ":"
		prefixLower := strings.ToLower(prefixUpper)
		// 如果匹配上了, 则说明规则中限定了请求方式
		if strings.HasPrefix(cleanRule, prefixUpper) || strings.HasPrefix(cleanRule, prefixLower) {
			// 请求方式是否匹配
			matchedMethod = v == method
			// 去除规则中的请求方式前缀
			cleanRule = strings.TrimLeft(strings.TrimLeft(cleanRule, prefixLower), prefixUpper)
			break
		}
	}

	// 非正则匹配
	if !strings.HasPrefix(rule, "^") {
		return matchedMethod && (strings.TrimLeft(cleanRule, "/") == url)
	}

	// 正则匹配
	re, err := regexp.Compile(cleanRule)
	if err == nil {
		return matchedMethod && re.MatchString(url)
	}

	return false
}
