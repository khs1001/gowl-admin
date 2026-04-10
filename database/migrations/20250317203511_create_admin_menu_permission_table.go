package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250317203511CreateAdminMenuPermissionTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250317203511CreateAdminMenuPermissionTable) Signature() string {
	return "20250317203511_create_admin_menu_permission_table"
}

// Up Run the migrations.
func (r *M20250317203511CreateAdminMenuPermissionTable) Up() error {
	if !facades.Schema().HasTable("admin_menu_permission") {
		return facades.Schema().Create("admin_menu_permission", func(table schema.Blueprint) {
			table.UnsignedBigInteger("admin_menu_id")
			table.UnsignedBigInteger("admin_permission_id")
			table.Unique("admin_menu_id", "admin_permission_id")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250317203511CreateAdminMenuPermissionTable) Down() error {
	return facades.Schema().DropIfExists("admin_menu_permission")
}
