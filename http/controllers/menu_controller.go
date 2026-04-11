package controllers

import (
	"github.com/khs1001/gowl-admin/models"
	"github.com/khs1001/gowl-admin/support/core"
)

type MenuController struct {
	*core.CrudController[models.AdminMenu]
}

func NewMenuController() *MenuController {
	return &MenuController{
		CrudController: core.NewCrudController[models.AdminMenu](),
	}
}
