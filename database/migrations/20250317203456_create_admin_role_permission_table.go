package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250317203456CreateAdminRolePermissionTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250317203456CreateAdminRolePermissionTable) Signature() string {
	return "20250317203456_create_admin_role_permission_table"
}

// Up Run the migrations.
func (r *M20250317203456CreateAdminRolePermissionTable) Up() error {
	if !facades.Schema().HasTable("admin_role_permission") {
		return facades.Schema().Create("admin_role_permission", func(table schema.Blueprint) {
			table.UnsignedBigInteger("admin_permission_id")
			table.UnsignedBigInteger("admin_role_id")

			table.Unique("admin_role_id", "admin_permission_id")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250317203456CreateAdminRolePermissionTable) Down() error {
	return facades.Schema().DropIfExists("admin_role_permission")
}
