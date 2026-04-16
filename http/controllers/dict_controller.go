package controllers

import (
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/contracts/http"
	"github.com/khs1001/gowl-admin/models"
	"github.com/khs1001/gowl-admin/services"
	"github.com/khs1001/gowl-admin/support/core"
)

type DictController struct {
	*core.CrudController[models.AdminDict]
}

func NewDictController() *DictController {
	c := &DictController{
		CrudController: core.NewCrudController[models.AdminDict](services.NewDictService()),
	}
	c.ListParams.IsTreeList = true
	c.ListParams.ParentIdField = "ParentValue"
	c.ListParams.IdField = "Value"
	c.ListParams.ChildrenField = "Children"
	c.ListParams.RootParentId = ""
	return c
}

func (c *DictController) DictOptions(ctx http.Context) http.Response {
	scopeSelect := func(query orm.Query) orm.Query {
		return query.Select("value", "label").
			Where("enabled", 1).
			OrderBy("sort", "asc")

	}
	items, _, err := c.Service.GetList(ctx, scopeSelect)

	if err != nil {
		return c.Error(ctx, err)
	}
	return c.Success(ctx, items)
}
