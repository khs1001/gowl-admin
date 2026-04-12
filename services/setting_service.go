package services

import (
	"context"

	"github.com/goravel/framework/facades"
	"github.com/khs1001/gowl-admin/models"
	"gorm.io/datatypes"
)

type SettingService struct {
}

func NewSettingService() *SettingService {
	return &SettingService{}
}

func (s *SettingService) Set(ctx context.Context, key string, value datatypes.JSONMap) (err error) {
	var adminSetting models.AdminSetting
	adminSetting.Key = key
	adminSetting.Value = value
	err = facades.Orm().WithContext(ctx).Query().Where("key", key).UpdateOrCreate(&adminSetting,
		map[string]interface{}{"key": adminSetting.Key},
		map[string]interface{}{"value": adminSetting.Value},
	)
	return err
}

func (s *SettingService) Get(ctx context.Context, key string) (result datatypes.JSONMap, err error) {
	var adminSetting models.AdminSetting
	err = facades.Orm().WithContext(ctx).Query().Where("key", key).First(&adminSetting)
	if err != nil {
		return nil, err
	}
	return adminSetting.Value, nil
}

func (s *SettingService) Scan(ctx context.Context, key string, pointer any, mapping ...map[string]string) (err error) {
	value, err := s.Get(ctx, key)
	if err != nil {
		return err
	}
	return value.Scan(pointer)
}
