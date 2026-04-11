package main

import (
	"os"

	"github.com/goravel/framework/packages"
	"github.com/goravel/framework/packages/match"
	"github.com/goravel/framework/packages/modify"
	"github.com/goravel/framework/support/file"
	"github.com/goravel/framework/support/path"
)

func main() {
	setup := packages.Setup(os.Args)
	serviceProvider := "&admin.ServiceProvider{}"
	moduleImport := setup.Paths().Module().Import()
	guard := `map[string]any{
		"driver":   "jwt",
		"provider": "user",
	}`
	authConfigPath := path.Config("auth.go")
	guardsConfig := match.Config("auth.guards")
	configPath := path.Config("admin.go")
	config, err := file.GetPackageContent(setup.Paths().Module().String(), "config/admin.go")
	if err != nil {
		panic(err)
	}

	setup.Install(
		// Register the service provider
		modify.GoFile(path.Bootstrap("providers.go")).
			Find(match.Imports()).Modify(modify.AddImport(moduleImport, "admin")).
			Find(match.Providers()).Modify(modify.Register(serviceProvider)),

		// Add config
		modify.File(configPath).Overwrite(config),
		// Add oss disk to filesystems.go
		modify.GoFile(authConfigPath).
			Find(guardsConfig).Modify(modify.AddConfig("admin", guard)),
	).Uninstall(
		// Remove auth guard
		modify.GoFile(authConfigPath).
			Find(guardsConfig).Modify(modify.RemoveConfig("admin")),

		// Remove config/admin.go
		modify.File(configPath).Remove(),

		// Unregister the service provider
		modify.GoFile(path.Bootstrap("providers.go")).
			Find(match.Providers()).Modify(modify.Unregister(serviceProvider)).
			Find(match.Imports()).Modify(modify.RemoveImport(moduleImport, "admin")),
	).Execute()
}
