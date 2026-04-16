package models

import (
	"encoding/json"

	"github.com/goravel/framework/database/orm"
)

// AdminDicts 表示 admin_dicts 表的模型
type AdminDict struct {
	orm.Model
	ParentValue string `gorm:"comment:上级" json:"parent_value"`
	// Type 字典类型，移除 size 属性，JSON 标签为 "type"
	Type string `gorm:"comment:字典类型" json:"type"`
	// Label 标签，移除 size 属性，JSON 标签为 "label"
	Label string `gorm:"comment:标签" json:"label"`
	// Value 字典值，移除 size 属性，JSON 标签为 "value"
	Value string `gorm:"comment:字典值" json:"value"`
	// Options 配置，JSON 标签为 "options"
	Options json.RawMessage `gorm:"type:json;comment:配置" json:"options"`
	// Enabled 可用状态，JSON 标签为 "enabled"
	Enabled *int8 `gorm:"type:tinyint;comment:可用;default:1" json:"enabled"`
	// Sort 排序，JSON 标签为 "sort"
	Sort *int8 `gorm:"type:tinyint;comment:排序;default:0" json:"sort"`
	// Remark 备注，移除 size 属性，JSON 标签为 "remark"
	Remark   string       `gorm:"comment:备注" json:"remark"`
	Children []*AdminDict `json:"children" gorm:"-"`
}
