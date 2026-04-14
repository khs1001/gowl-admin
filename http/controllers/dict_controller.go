package controllers

import (
	"github.com/khs1001/gowl-admin/models"
	"github.com/khs1001/gowl-admin/support/core"
)

type DictController struct {
	*core.CrudController[models.AdminDict]
}

func NewDictController() *DictController {
	return &DictController{
		CrudController: core.NewCrudController[models.AdminDict](),
	}
}
