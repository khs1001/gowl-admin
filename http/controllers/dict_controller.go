package controllers

import (
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
