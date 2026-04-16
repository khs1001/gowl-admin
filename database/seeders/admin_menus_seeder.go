package seeders

import (
	"github.com/khs1001/gowl-admin/models"

	"github.com/goravel/framework/facades"
)

type AdminMenusSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *AdminMenusSeeder) Signature() string {
	return "AdminMenusSeeder"
}

// Run executes the seeder logic.
func (s *AdminMenusSeeder) Run() error {
	data := []*models.AdminMenu{
		{
			Title:   "控制台",
			Icon:    "mdi:chart-line",
			Url:     "/dashboard",
			UrlType: 1,
			IsHome:  1,
			Visible: 1,
		},
	}
	err := facades.Orm().Query().Create(data)
	if err != nil {
		return err
	}
	authData := models.AdminMenu{
		Title:   "权限管理",
		Icon:    "icon-park-outline:people-safe",
		Url:     "/auth",
		Visible: 1,
		UrlType: 1,
	}
	err = facades.Orm().Query().Create(&authData)
	if err != nil {
		return err
	}
	data = []*models.AdminMenu{
		{
			ParentId: int(authData.ID),
			Title:    "管理员",
			Icon:     "ph:user-gear",
			Url:      "/auth/admin_users",
			Visible:  1,
			UrlType:  1,
		},
		{
			ParentId: int(authData.ID),
			Title:    "角色",
			Icon:     "carbon:user-role",
			Url:      "/auth/admin_roles",
			Visible:  1,
			UrlType:  1,
		},
		{
			ParentId: int(authData.ID),
			Title:    "菜单",
			Icon:     "ant-design:menu-unfold-outlined",
			Url:      "/auth/admin_menus",
			Visible:  1,
			UrlType:  1,
		},
	}
	err = facades.Orm().Query().Create(data)
	if err != nil {
		return err
	}
	systemData := models.AdminMenu{
		Title:   "系统管理",
		Icon:    "material-symbols:settings-outline",
		Url:     "/system",
		Visible: 1,
		UrlType: 1,
	}
	err = facades.Orm().Query().Create(&systemData)
	if err != nil {
		return err
	}
	data = []*models.AdminMenu{
		{
			ParentId: int(systemData.ID),
			Title:    "数据字典",
			Icon:     "fluent-mdl2:dictionary",
			Url:      "/system/admin_dicts",
			Visible:  1,
			UrlType:  1,
		},
		{
			ParentId: int(systemData.ID),
			Title:    "Api服务商",
			Icon:     "streamline-freehand:server-api-cloud",
			Url:      "/system/admin_api_providers",
			Visible:  1,
			UrlType:  1,
		},
	}
	err = facades.Orm().Query().Create(data)
	if err != nil {
		return err
	}
	return nil
}
