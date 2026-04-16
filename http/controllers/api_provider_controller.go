package controllers

import (
	"github.com/khs1001/gowl-admin/models"
	"github.com/khs1001/gowl-admin/support/core"
)

type ApiProviderController struct {
	*core.CrudController[models.AdminDict]
}

func NewApiProviderController() *ApiProviderController {
	return &ApiProviderController{
		CrudController: core.NewCrudController[models.AdminDict](),
	}
}
