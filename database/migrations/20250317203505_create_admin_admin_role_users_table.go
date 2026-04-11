package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250317203505CreateAdminRoleUsersTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250317203505CreateAdminRoleUsersTable) Signature() string {
	return "20250317203505_create_admin_role_users_table"
}

// Up Run the migrations.
func (r *M20250317203505CreateAdminRoleUsersTable) Up() error {
	if !facades.Schema().HasTable("admin_role_users") {
		return facades.Schema().Create("admin_role_users", func(table schema.Blueprint) {
			table.UnsignedBigInteger("role_id")
			table.UnsignedBigInteger("user_id")
			table.Timestamps()
			table.Unique("role_id", "user_id")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250317203505CreateAdminRoleUsersTable) Down() error {
	return facades.Schema().DropIfExists("admin_role_users")
}
