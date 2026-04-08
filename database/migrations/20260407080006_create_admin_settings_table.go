package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"github.com/goravel/framework/facades"
)

type M20260407080006CreateAdminSettingsTable struct{}

// Signature The unique signature for the migration.
func (r *M20260407080006CreateAdminSettingsTable) Signature() string {
	return "20260407080006_create_admin_settings_table"
}

// Up Run the migrations.
func (r *M20260407080006CreateAdminSettingsTable) Up() error {
	if !facades.Schema().HasTable("admin_settings") {
		return facades.Schema().Create("admin_settings", func(table schema.Blueprint) {
			table.String("key").Comment("键")
			table.Jsonb("values").Comment("值")
			table.TimestampsTz()
			table.Primary("key")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20260407080006CreateAdminSettingsTable) Down() error {
	return facades.Schema().DropIfExists("admin_settings")
}
