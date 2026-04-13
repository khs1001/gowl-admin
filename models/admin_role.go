package models

import (
	"github.com/goravel/framework/database/orm"
)

type AdminRole struct {
	orm.Model
	Name string `json:"name" db:"name"`
	Slug string `json:"slug" db:"slug"`
}

func (r *AdminRole) TableName() string {
	return "admin_roles"
}
