package models

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/goravel/framework/database/orm"
)

type AdminSetting struct {
	orm.Model
	Key   string      `gorm:"type:varchar(50);unique;comment:键名"`
	Value *gjson.Json `gorm:"type:json;comment:键值"`
}
