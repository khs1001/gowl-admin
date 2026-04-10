package models

import (
	"github.com/goravel/framework/database/orm"
)

type AdminUser struct {
	orm.Model
	Username      string       `json:"username" db:"username"`
	Password      string       `json:"-" db:"password"`
	Enabled       int8         `json:"enabled" db:"enabled"`
	Name          string       `json:"name" db:"name"`
	Avatar        string       `json:"avatar" db:"avatar"`
	RememberToken string       `json:"remember_token" db:"remember_token"`
	Roles         []*AdminRole `gorm:"many2many:admin_role_users;joinForeignKey:user_id;joinReferences:role_id" json:"roles"`
}

func (r *AdminUser) TableName() string {
	return "admin_users"
}
