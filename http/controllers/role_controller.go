package controllers

import (
	"github.com/khs1001/gowl-admin/models"
	"github.com/khs1001/gowl-admin/support/core"
)

type RoleController struct {
	*core.CrudController[models.AdminRole]
}

func NewRoleController() *RoleController {
	return &RoleController{
		CrudController: core.NewCrudController[models.AdminRole](),
	}
}
