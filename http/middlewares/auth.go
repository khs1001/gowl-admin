package middlewares

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"github.com/khs1001/gowl-admin/consts"
	"github.com/khs1001/gowl-admin/support/core"
)

// Auth 管理后台认证中间件
// 负责验证管理后台请求的身份认证，包括请求类型检查、验证开关、白名单校验和Token验证
func Auth(ctx http.Context) {
	facades.Log().Debug(ctx.Request().Path())
	// 非后台请求, 跳过
	if !core.IsAdminRequest(ctx.Request().Path()) {
		ctx.Request().Next()
		return
	}
	// 未开启验证
	if !facades.Config().GetBool(consts.AdminAuthEnabled, false) {
		ctx.Request().Next()
		return
	}

	// 白名单
	exclude := facades.Config().Get(consts.AdminAuthExclude, []string{}).([]string)
	if len(exclude) > 0 {
		for _, v := range exclude {
			if core.IsAllowRequest(v, ctx.Request().Method(), ctx.Request().Path()) {
				ctx.Request().Next()
				return
			}
		}
	}
	token := ctx.Request().Header(consts.Authorization)

	facades.Log().Debug("token", token)
	if token == "" {
		core.UnAuthorized(ctx)
		return
	}

	guard := facades.Config().GetString(consts.AdminAuthGuard)
	payload, err := facades.Auth(ctx).Guard(guard).Parse(token)
	facades.Log().Debugf("payload %v, err %v", payload, err)

	if err != nil {
		core.UnAuthorized(ctx)
		return
	}
	// 将用户ID添加到上下文中，供后续请求处理使用
	ctx.WithValue(consts.UserID, payload.Key)
	ctx.Request().Next()
}
