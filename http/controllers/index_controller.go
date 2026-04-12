package controllers

import (
	"path/filepath"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/khs1001/gowl-admin/consts"
	"github.com/khs1001/gowl-admin/http/responses"
	"github.com/khs1001/gowl-admin/services"
	"github.com/khs1001/gowl-admin/support/core"
	"github.com/khs1001/gowl-admin/support/schema"
)

type IndexController struct {
	*core.BaseController
	SettingService *services.SettingService
}

func NewIndexController() *IndexController {
	return &IndexController{
		SettingService: services.NewSettingService(),
	}
}

func (c *IndexController) Settings(ctx http.Context) http.Response {
	var setting responses.AdminTheme
	err := gconv.Struct(facades.Config().Get(consts.Admin), &setting)
	if err != nil {
		return c.Error(ctx, err)
	}
	//从数据库里读取系统配置
	if setting.SystemThemeSetting, err = c.SettingService.Get(ctx, ctx.Request().Path()); err != nil {
		return c.Error(ctx, err)
	}
	setting.Locale = facades.Config().GetString(consts.AppLocal)
	setting.SetLocaleOptions()
	return c.Success(ctx, setting)
}

func (c *IndexController) SaveSettings(ctx http.Context) http.Response {
	var setting responses.AdminTheme
	err := ctx.Request().Bind(&setting)
	if err != nil {
		return c.Error(ctx, err)
	}
	err = c.SettingService.Set(ctx, ctx.Request().Path(), setting.SystemThemeSetting)
	if err != nil {
		return c.Error(ctx, err)
	}
	return c.Success(ctx, nil)
}

func (c *IndexController) PageSchema(ctx http.Context) http.Response {
	sign := ctx.Request().Query(consts.Sign)
	path := ctx.Request().Query(consts.Path)
	if path == "" {
		path = gstr.Replace(filepath.Dir(sign), "\\", "/")
		sign = filepath.Base(sign)
	}
	schema := schema.GetPage(ctx, sign, path, gconv.Map(ctx.Request().All()))
	return c.Success(ctx, schema)
}
