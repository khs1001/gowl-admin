package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250319132111CreateAdminDictsTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250319132111CreateAdminDictsTable) Signature() string {
	return "20250319132111_create_admin_dicts_table"
}

// Up Run the migrations.
func (r *M20250319132111CreateAdminDictsTable) Up() error {
	if !facades.Schema().HasTable("admin_dicts") {
		return facades.Schema().Create("admin_dicts", func(table schema.Blueprint) {
			table.ID()
			table.String("type", 50).Comment("字典类型")
			table.String("label", 50).Comment("标签")
			table.String("value", 50).Comment("字典值")
			table.Json("options").Nullable().Comment("配置")
			table.TinyInteger("enabled").Nullable().Comment("可用").Default(1)
			table.TinyInteger("sort").Nullable().Comment("排序").Default(0)
			table.String("remark", 255).Nullable().Comment("备注")
			table.Timestamps()

			table.Unique("type", "value")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250319132111CreateAdminDictsTable) Down() error {
	return facades.Schema().DropIfExists("admin_dicts")
}
