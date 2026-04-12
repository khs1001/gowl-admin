package models

import (
	"github.com/goravel/framework/database/orm"
	"gorm.io/datatypes"
)

type AdminSetting struct {
	orm.Model
	Key   string            `gorm:"type:varchar(50);unique;comment:键名"`
	Value datatypes.JSONMap `gorm:"type:json;comment:键值"`
}
