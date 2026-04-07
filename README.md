# admin

## 安装
```
./artisan package:install github.com/khs1001/gowl-admin
```

## 卸载
```
./artisan package:uninstall github.com/khs1001/gowl-admin
```

## 发资资源
```
./artisan vendor:publish --package=github.com/khs1001/gowl-admin  --force
```

## 生成模型
```
go run . artisan make:model --table=admin_settings -f AdminSetting
go run . artisan make:model --table=admin_menus -f AdminMenu
go run . artisan make:model --table=admin_permissions -f AdminPermission
go run . artisan make:model --table=admin_users -f AdminUser
go run . artisan make:model --table=admin_roles -f AdminRole
go run . artisan make:model --table=admin_role_permissions -f AdminRolePermission
```
