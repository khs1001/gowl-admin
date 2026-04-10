package models

import (
	"github.com/goravel/framework/database/orm"
)

type AdminRolePermission struct {
	orm.Timestamps
	RoleId       int `json:"role_id" db:"role_id"`
	PermissionId int `json:"permission_id" db:"permission_id"`
}

func (r *AdminRolePermission) TableName() string {
	return "admin_role_permissions"
}
