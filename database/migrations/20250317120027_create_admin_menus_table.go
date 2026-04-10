package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250317120027CreateAdminMenusTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250317120027CreateAdminMenusTable) Signature() string {
	return "20250317120027_create_admin_menus_table"
}

// Up Run the migrations.
func (r *M20250317120027CreateAdminMenusTable) Up() error {
	if !facades.Schema().HasTable("admin_menus") {
		return facades.Schema().Create("admin_menus", func(table schema.Blueprint) {
			table.ID()
			table.UnsignedBigInteger("parent_id").Comment("父级ID").Default(0)
			table.String("name", 50).Comment("菜单名称")
			table.String("icon", 50).Comment("菜单图标")
			table.String("path", 255).Comment("菜单路径")
			table.TinyInteger("visible").Comment("是否可见").Default(1)
			table.Integer("sort").Comment("排序").Default(0)
			table.TinyInteger("is_home").Comment("是否首页").Default(0)
			table.TinyInteger("is_full").Comment("是否全屏").Default(0)
			table.TinyInteger("keep_alive").Comment("是否缓存").Default(0)
			table.String("page_sign", 255).Comment("页面标识")
			table.Timestamps()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250317120027CreateAdminMenusTable) Down() error {
	return facades.Schema().DropIfExists("admin_menus")
}
