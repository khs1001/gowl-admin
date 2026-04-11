package seeders

import (
	"github.com/khs1001/gowl-admin/models"

	"github.com/goravel/framework/facades"
)

type AdminRolesSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *AdminRolesSeeder) Signature() string {
	return "AdminRolesSeeder"
}

// Run executes the seeder logic.
func (s *AdminRolesSeeder) Run() error {
	data := []*models.AdminRole{
		{
			Name: "系统管理员",
			Slug: "administrator",
		},
	}
	return facades.Orm().Query().Create(data)
}
