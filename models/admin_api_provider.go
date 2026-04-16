package models

import (
	"encoding/json"

	"github.com/goravel/framework/database/orm"
)

// AdminApiProvider
type AdminApiProvider struct {
	ID           uint            `gorm:"primaryKey" json:"id"`
	Type         string          `gorm:"comment:类型" json:"type"`
	Protocol     string          `gorm:"comment:协议" json:"protocol"`
	Name         string          `gorm:"comment:名称" json:"name"`
	Code         string          `gorm:"comment:标识" json:"code"`
	Options      json.RawMessage `gorm:"type:json;comment:配置" json:"options"`
	Sort         *int8           `gorm:"type:tinyint;comment:排序;default:0" json:"sort"`
	Remark       string          `gorm:"comment:备注" json:"remark"`
	Children     []*AdminDict    `json:"children" gorm:"-"`
	TypeInfo     *AdminDict      `json:"type_info" gorm:"hasOne;foreignKey:Value;references:Type;where:type='api_type'"`
	ProtocolInfo *AdminDict      `json:"protocol_info" gorm:"hasOne;foreignKey:Value;references:Type;where:type='api_protocol'"`
	orm.Timestamps
}
