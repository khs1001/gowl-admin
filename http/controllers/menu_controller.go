package controllers

import (
	"github.com/khs1001/gowl-admin/models"
	"github.com/khs1001/gowl-admin/support/core"
)

type MenuController struct {
	*core.CrudController[models.AdminMenu]
}

func NewMenuController() *MenuController {
	c := &MenuController{
		CrudController: core.NewCrudController[models.AdminMenu](),
	}
	c.ListParams.IsTreeList = true
	c.ListParams.ChildrenField = "Children"
	c.ListParams.ParentIdField = "ParentId"
	c.ListParams.IdField = "ID"
	c.ListParams.RootParentId = 0

	return c

}
