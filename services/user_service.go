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
	s := &UserService{
		CrudService: core.NewCrudService[models.AdminUser](),
	}
	s.SetWiths([]string{"Roles"})
	s.SetAssociations([]string{"Roles"})

	return s
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
