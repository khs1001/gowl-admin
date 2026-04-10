package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250317112835CreateAdminSettingsTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250317112835CreateAdminSettingsTable) Signature() string {
	return "20250317112835_create_admin_settings_table"
}

// Up Run the migrations.
func (r *M20250317112835CreateAdminSettingsTable) Up() error {
	if !facades.Schema().HasTable("admin_settings") {
		return facades.Schema().Create("admin_settings", func(table schema.Blueprint) {
			table.ID()
			table.String("key", 50).Comment("键名")
			table.Json("value").Comment("键值")
			table.Timestamps()

			table.Unique("key")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250317112835CreateAdminSettingsTable) Down() error {
	return facades.Schema().DropIfExists("admin_settings")
}
