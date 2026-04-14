package controllers

import (
	"github.com/khs1001/gowl-admin/models"
	"github.com/khs1001/gowl-admin/services"
	"github.com/khs1001/gowl-admin/support/core"
)

type UserController struct {
	*core.CrudController[models.AdminUser]
}

func NewUserController() *UserController {
	c := &UserController{
		CrudController: core.NewCrudController[models.AdminUser](services.NewUserService()),
	}
	return c
}
