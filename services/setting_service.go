package services

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/goravel/framework/facades"
	"github.com/khs1001/gowl-admin/models"
)

type SettingService struct {
}

func NewSettingService() *SettingService {
	return &SettingService{}
}

func (s *SettingService) Set(ctx context.Context, key string, value any) (err error) {
	var adminSetting models.AdminSetting
	adminSetting.Key = key
	adminSetting.Value = gjson.New(value)
	_, err = facades.Orm().WithContext(ctx).Query().Where("key", key).Update(&adminSetting)
	return err
}

func (s *SettingService) Get(ctx context.Context, key string) (result *gjson.Json, err error) {
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
	return value.Scan(pointer, mapping...)
}
