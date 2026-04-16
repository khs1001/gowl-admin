package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"github.com/goravel/framework/facades"
)

type M20260414143735CreateAdminApiProvidersTable struct{}

// Signature The unique signature for the migration.
func (r *M20260414143735CreateAdminApiProvidersTable) Signature() string {
	return "20260414143735_create_admin_api_providers_table"
}

// Up Run the migrations.
func (r *M20260414143735CreateAdminApiProvidersTable) Up() error {
	if !facades.Schema().HasTable("admin_api_providers") {
		return facades.Schema().Create("admin_api_providers", func(table schema.Blueprint) {
			table.ID()
			table.String("type", 50).Comment("类型")
			table.String("protocol", 50).Comment("协议")
			table.String("name", 50).Comment("名称")
			table.String("code", 50).Comment("标识")
			table.TinyInteger("Enabled").Comment("启用")
			table.Jsonb("options").Comment("参数")
			table.Text("remark").Comment("备注")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20260414143735CreateAdminApiProvidersTable) Down() error {
	return facades.Schema().DropIfExists("admin_api_providers")
}
