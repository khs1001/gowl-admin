package controllers

import (
	"github.com/khs1001/gowl-admin/models"
	"github.com/khs1001/gowl-admin/support/core"
)

type UserController struct {
	*core.CrudController[models.AdminUser]
}

func NewUserController() *UserController {
	return &UserController{
		CrudController: core.NewCrudController[models.AdminUser](),
	}
}
