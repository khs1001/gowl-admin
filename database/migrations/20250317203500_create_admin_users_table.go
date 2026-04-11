package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/contracts/database/seeder"
	"github.com/goravel/framework/facades"
	"github.com/khs1001/gowl-admin/database/seeders"
)

type M20250317203500CreateAdminUsersTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250317203500CreateAdminUsersTable) Signature() string {
	return "20250317203500_create_admin_users_table"
}

// Up Run the migrations.
func (r *M20250317203500CreateAdminUsersTable) Up() error {
	if !facades.Schema().HasTable("admin_users") {
		err := facades.Schema().Create("admin_users", func(table schema.Blueprint) {
			table.ID()
			table.String("username", 50).Comment("用户名")
			table.String("password", 255).Comment("密码")
			table.String("name", 50).Comment("名称")
			table.String("avatar", 255).Comment("头像")
			table.TinyInteger("enabled").Default(1).Comment("是否启用")
			table.String("remember_token")
			table.Timestamps()
			table.Unique("username")
		})
		if err != nil {
			return err
		}
		facades.Seeder().CallOnce([]seeder.Seeder{
			&seeders.AdminUsersSeeder{},
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250317203500CreateAdminUsersTable) Down() error {
	return facades.Schema().DropIfExists("admin_users")
}
