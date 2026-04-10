package seeders

import (
	"github.com/khs1001/gowl-admin/models"

	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
)

type AdminUsersSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *AdminUsersSeeder) Signature() string {
	return "AdminUsersSeeder"
}

// Run executes the seeder logic.
func (s *AdminUsersSeeder) Run() error {
	password, _ := facades.Hash().Make("admin")
	data := []*models.AdminUser{
		{
			Username: "admin",
			Name:     "Administrator",
			Password: password,
			Enabled:  1,
			Avatar:   "/admin-assets/default-avatar.png",
			Roles: []*models.AdminRole{
				{
					Model: orm.Model{ID: 1},
				},
			},
		},
	}
	return facades.Orm().Query().Select(orm.Associations).Create(data)
}
