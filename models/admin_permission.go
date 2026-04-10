package models

import (
	"github.com/goravel/framework/database/orm"
)

type AdminPermission struct {
	orm.Model
	Name        string             `json:"name" db:"name"`
	Slug        string             `json:"slug" db:"slug"`
	HttpMethod  string             `json:"http_method" db:"http_method"`
	HttpPath    string             `json:"http_path" db:"http_path"`
	CustomOrder int                `json:"custom_order" db:"custom_order"`
	ParentId    int                `json:"parent_id" db:"parent_id"`
	Children    []*AdminPermission `json:"children" gorm:"-"`
}

func (r *AdminPermission) TableName() string {
	return "admin_permissions"
}
