package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250317203505CreateAdminUserRoleTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250317203505CreateAdminUserRoleTable) Signature() string {
	return "20250317203505_create_admin_user_role_table"
}

// Up Run the migrations.
func (r *M20250317203505CreateAdminUserRoleTable) Up() error {
	if !facades.Schema().HasTable("admin_user_role") {
		return facades.Schema().Create("admin_user_role", func(table schema.Blueprint) {
			table.UnsignedBigInteger("admin_user_id")
			table.UnsignedBigInteger("admin_role_id")
			table.Unique("admin_user_id", "admin_role_id")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250317203505CreateAdminUserRoleTable) Down() error {
	return facades.Schema().DropIfExists("admin_user_role")
}
