package controllers

import (
	"errors"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/khs1001/gowl-admin/consts"
	"github.com/khs1001/gowl-admin/models"
	"github.com/khs1001/gowl-admin/services"
	"github.com/khs1001/gowl-admin/support/core"
	"github.com/khs1001/gowl-admin/support/schema"
)

type AuthController struct {
	*core.BaseController
	AuthService *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		AuthService: services.NewAuthService(),
	}
}

// Login 登录
func (c *AuthController) Login(ctx http.Context) http.Response {
	validator, err := ctx.Request().Validate(map[string]string{
		consts.Username: "required|max_len:32",
		consts.Password: "required|min_len:5|max_len:32",
	})
	if err != nil {
		return c.Error(ctx, err)
	}
	if validator.Fails() {
		err = errors.New(validator.Errors().One())
		return c.Error(ctx, err)
	}
	token, err := c.AuthService.Login(ctx,
		ctx.Request().Input(consts.Username),
		ctx.Request().Input(consts.Password))
	if err != nil {
		return c.Error(ctx, err)
	}
	return c.Success(ctx, &http.Json{
		"token": token,
	})
}

// Logout 退出登录
func (c *AuthController) Logout(ctx http.Context) http.Response {
	return c.Success(ctx, http.Json{})
}

// CurrentUser 获取当前用户信息
func (c *AuthController) CurrentUser(ctx http.Context) http.Response {
	var user models.AdminUser
	err := facades.Orm().WithContext(ctx).Query().
		Where(consts.ID, core.UserID(ctx)).
		FindOrFail(&user)
	if err != nil {
		return c.Error(ctx, err)
	}
	menus := schema.GetPage(ctx, "user_menu#button", ctx.Request().Path(), gconv.Map(user))
	return c.Success(ctx, &http.Json{
		"name":   user.Name,
		"avatar": user.Avatar,
		"menus":  menus,
	})
}

// Menus 获取菜单信息
func (c *AuthController) Menus(ctx http.Context) http.Response {
	var items []*models.AdminMenu
	err := facades.Orm().WithContext(ctx).Query().
		OrderBy(consts.CustomOrder).Get(&items)
	if err != nil {
		return c.Error(ctx, err)
	}
	menus := c.AuthService.BuildRoutes(ctx, items, 0)
	return c.Success(ctx, menus)
}
