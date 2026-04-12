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
		{
			Title:   "系统管理",
			Icon:    "material-symbols:settings-outline",
			Url:     "/system",
			Visible: 1,
			UrlType: 1,
		},
		{
			ParentId: 2,
			Title:    "管理员",
			Icon:     "ph:user-gear",
			Url:      "/system/admin_users",
			Visible:  1,
			UrlType:  1,
		},
		{
			ParentId: 2,
			Title:    "角色",
			Icon:     "carbon:user-role",
			Url:      "/system/admin_roles",
			Visible:  1,
			UrlType:  1,
		},
		{
			ParentId: 2,
			Title:    "菜单",
			Icon:     "ant-design:menu-unfold-outlined",
			Url:      "/system/admin_menus",
			Visible:  1,
			UrlType:  1,
		},
	}
	return facades.Orm().Query().Create(data)
}
