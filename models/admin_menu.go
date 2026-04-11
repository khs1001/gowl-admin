package models

import (
	"github.com/goravel/framework/database/orm"
)

type AdminMenu struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	ParentId    int          `json:"parent_id" db:"parent_id"`
	CustomOrder int          `json:"custom_order" db:"custom_order"`
	Title       string       `json:"title" db:"title"`
	Icon        string       `json:"icon" db:"icon"`
	Url         string       `json:"url" db:"url"`
	UrlType     int8         `json:"url_type" db:"url_type"`
	Visible     int8         `json:"visible" db:"visible"`
	IsHome      int8         `json:"is_home" db:"is_home"`
	KeepAlive   int8         `json:"keep_alive" db:"keep_alive"`
	IframeUrl   string       `json:"iframe_url" db:"iframe_url"`
	Component   string       `json:"component" db:"component"`
	IsFull      int8         `json:"is_full" db:"is_full"`
	Extension   string       `json:"extension" db:"extension"`
	Children    []*AdminMenu `json:"children" gorm:"-"`
	orm.Timestamps
}

func (r *AdminMenu) TableName() string {
	return "admin_menus"
}
