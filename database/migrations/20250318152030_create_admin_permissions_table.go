package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

// M20250318152030CreateAdminPermissionsTable 该迁移用于创建 admin_permissions 表
type M20250318152030CreateAdminPermissionsTable struct {
}

// Signature 返回迁移的唯一签名
func (r *M20250318152030CreateAdminPermissionsTable) Signature() string {
	return "20250318152030_create_admin_permissions_table"
}

// Up 执行迁移操作，创建 admin_permissions 表
func (r *M20250318152030CreateAdminPermissionsTable) Up() error {
	if !facades.Schema().HasTable("admin_permissions") {
		return facades.Schema().Create("admin_permissions", func(table schema.Blueprint) {
			table.ID()
			table.UnsignedBigInteger("parent_id").Default(0).Comment("父级ID")
			table.String("name", 255).Comment("名称")
			table.String("sign", 255).Comment("标识")
			table.Json("api").Comment("api")
			table.Integer("sort").Default(0).Comment("排序")
			table.Timestamps()
			table.Unique("sign")
		})
	}

	return nil
}

// Down 回滚迁移操作，删除 admin_permissions 表
func (r *M20250318152030CreateAdminPermissionsTable) Down() error {
	return facades.Schema().DropIfExists("admin_permissions")
}
