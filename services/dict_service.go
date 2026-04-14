package services

import (
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/contracts/http"
	"github.com/khs1001/gowl-admin/models"
	"github.com/khs1001/gowl-admin/support/core"
)

type DictService struct {
	*core.CrudService[models.AdminDict]
}

func NewDictService() *DictService {
	s := &DictService{
		CrudService: core.NewCrudService[models.AdminDict](),
	}

	return s
}

func (s *DictService) GetList(ctx http.Context, scopes ...func(orm.Query) orm.Query) (items any, total int64, err error) {
	scopeWhere := func(query orm.Query) orm.Query {
		return query.Where("type", ctx.Request().Input("type"))
	}
	scopes = append(scopes, scopeWhere)
	return s.CrudService.GetList(ctx, scopes...)
}
