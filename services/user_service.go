package services

import (
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/khs1001/gowl-admin/models"
	"github.com/khs1001/gowl-admin/support/core"
)

type UserService struct {
	*core.CrudService[models.AdminUser]
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Create(ctx http.Context, item any, scopes ...func(orm.Query) orm.Query) (err error) {
	//密码加密
	//初始把账号名当作密码	//密码加密
	if ctx.Request().Input("password") != "" {
		item.(*models.AdminUser).Password, _ = facades.Hash().Make(ctx.Request().Input("password"))
	}
	return s.CrudService.Create(ctx, item, scopes...)
}

func (s *UserService) Update(ctx http.Context, id any, item any, scopes ...func(orm.Query) orm.Query) (rowsAffected int64, err error) {
	//密码加密
	if ctx.Request().Input("password") != "" {
		item.(*models.AdminUser).Password, _ = facades.Hash().Make(ctx.Request().Input("password"))
	}
	return s.CrudService.Update(ctx, id, item, scopes...)
}

func (s *UserService) GetList(ctx http.Context, scopes ...func(orm.Query) orm.Query) (items any, total int64, err error) {
	//增加查询条件
	scopeWhere := func(query orm.Query) orm.Query {
		query = query.Where("type", ctx.Request().Input("type"))
		return query
	}
	scopes = append(scopes, scopeWhere)
	return s.CrudService.GetList(ctx, scopes...)
}
