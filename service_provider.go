package admin

import (
	"github.com/goravel/framework/contracts/binding"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/khs1001/gowl-admin/database/migrations"
	"github.com/khs1001/gowl-admin/http/middlewares"
	"github.com/khs1001/gowl-admin/routes"
)

const Binding = "admin"
const PackageName = "github.com/khs1001/gowl-admin"

var App foundation.Application

type ServiceProvider struct {
}

// Relationship returns the relationship of the service provider.
func (r *ServiceProvider) Relationship() binding.Relationship {
	return binding.Relationship{
		Bindings:     []string{},
		Dependencies: []string{},
		ProvideFor:   []string{},
	}
}

// Register registers the service provider.
func (r *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return &Admin{}, nil
	})
}

// Boot boots the service provider, will be called after all service providers are registered.
func (r *ServiceProvider) Boot(app foundation.Application) {
	r.RegisterPublishes(app)
	r.RegisterMigrations(app)
	// 注册管理员路由组
	app.MakeRoute().GlobalMiddleware(middlewares.Auth)
	app.MakeRoute().Group(routes.Admin)
}

// 注册资源发布
func (r *ServiceProvider) RegisterPublishes(app foundation.Application) {
	// 发布视图资源到应用的公共目录
	app.Publishes(PackageName, map[string]string{
		"admin-views/dist": app.PublicPath("admin-assets"),
		"resources":        app.ResourcePath(""),
	}, "views")
}

// 注册数据库迁移
func (r *ServiceProvider) RegisterMigrations(app foundation.Application) {
	//数据数据库迁移
	AllMigrations := app.MakeSchema().Migrations()
	AllMigrations = append(AllMigrations,
		&migrations.M20250317112835CreateAdminSettingsTable{},
		&migrations.M20250317203511CreateAdminMenuPermissionTable{},
		&migrations.M20250317203456CreateAdminRolePermissionTable{},
		&migrations.M20250317203505CreateAdminRoleUsersTable{},
		&migrations.M20250317120027CreateAdminMenusTable{},
		&migrations.M20250318152030CreateAdminPermissionsTable{},
		&migrations.M20250317203452CreateAdminRolesTable{},
		&migrations.M20250317203500CreateAdminUsersTable{},
		&migrations.M20250319132111CreateAdminDictsTable{},
	)
	app.MakeSchema().Register(AllMigrations)
}
