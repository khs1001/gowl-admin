package seeders

import (
	"github.com/khs1001/gowl-admin/models"

	"github.com/goravel/framework/database/orm"
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
			Title:  "控制台",
			Icon:   "mdi:chart-line",
			Url:    "/dashboard",
			IsHome: 1,
		},
		{
			Title: "系统管理",
			Icon:  "material-symbols:settings-outline",
			Url:   "/system",
		},
		{
			Title: "管理员",
			Icon:  "ph:user-gear",
			Url:   "/system/admin_users",
		},
		{
			Title: "角色",
			Icon:  "carbon:user-role",
			Url:   "/system/admin_roles",
		},
		{
			Title: "菜单",
			Icon:  "ant-design:menu-unfold-outlined",
			Url:   "/system/admin_menus",
		},
		{
			Title: "设置",
			Icon:  "akar-icons:settings-horizontal",
			Url:   "/system/settings",
		},
	}
	return facades.Orm().Query().Select(orm.Associations).Create(data)
}
