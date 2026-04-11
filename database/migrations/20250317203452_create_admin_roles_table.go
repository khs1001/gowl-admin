package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250317203452CreateAdminRolesTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250317203452CreateAdminRolesTable) Signature() string {
	return "20250317203452_create_admin_roles_table"
}

// Up Run the migrations.
func (r *M20250317203452CreateAdminRolesTable) Up() error {
	if !facades.Schema().HasTable("admin_roles") {
		return facades.Schema().Create("admin_roles", func(table schema.Blueprint) {
			table.ID()
			table.String("name", 50).Comment("角色名称")
			table.String("slug", 50).Comment("角色标识")
			table.Timestamps()
			table.Unique("slug")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250317203452CreateAdminRolesTable) Down() error {
	return facades.Schema().DropIfExists("admin_roles")
}
