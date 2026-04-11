package core

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/khs1001/gowl-admin/consts"
)

type ICrudService interface {
	Query(ctx http.Context) orm.Query
	GetList(ctx http.Context, scopes ...func(orm.Query) orm.Query) (items any, err error)
	GetPageList(ctx http.Context, page, perPage int, scopes ...func(orm.Query) orm.Query) (items any, total int64, err error)
	GetListScopes(ctx http.Context) []func(orm.Query) orm.Query
	GetDetail(ctx http.Context, id any, scopes ...func(orm.Query) orm.Query) (item any, err error)
	GetDetailScopes(ctx http.Context) []func(orm.Query) orm.Query
	Create(ctx http.Context, item any, scopes ...func(orm.Query) orm.Query) (err error)
	GetCreateScopes(ctx http.Context) []func(orm.Query) orm.Query
	GetCreateData(ctx http.Context) any
	Update(ctx http.Context, id any, item any, scopes ...func(orm.Query) orm.Query) (rowsAffected int64, err error)
	GetUpdateScopes(ctx http.Context) []func(orm.Query) orm.Query
	Delete(ctx http.Context, id any, scopes ...func(orm.Query) orm.Query) (rowsAffected int64, err error)
	GetDeleteScopes(ctx http.Context) []func(orm.Query) orm.Query
}

// CrudService 通用增删改查服务
type CrudService[Model any] struct {
}

func NewCrudService[Model any]() *CrudService[Model] {
	return &CrudService[Model]{}
}

func (c *CrudService[Model]) Query(ctx http.Context) orm.Query {
	var model Model
	return facades.Orm().WithContext(ctx).Query().Model(&model)
}

func (c *CrudService[Model]) GetListScopes(ctx http.Context) []func(orm.Query) orm.Query {
	return []func(orm.Query) orm.Query{
		c.ScpoeOrderBy(ctx),
		c.ScpoeWith(ctx),
	}
}
func (c *CrudService[Model]) GetList(ctx http.Context, scopes ...func(orm.Query) orm.Query) (items any, err error) {
	items = make([]*Model, 0)
	query := c.Query(ctx).Scopes(scopes...)
	err = query.Get(&items)
	return items, err
}

// GetPageList 分页查询
func (c *CrudService[Model]) GetPageList(ctx http.Context, page, perPage int, scopes ...func(orm.Query) orm.Query) (items any, total int64, err error) {
	items = make([]*Model, 0)
	query := c.Query(ctx).Scopes(scopes...)
	err = query.Paginate(page, perPage, &items, &total)
	return items, total, err
}

// ScpoeOrderBy 排序
func (c *CrudService[Model]) ScpoeOrderBy(ctx http.Context) func(orm.Query) orm.Query {
	orderBy := ctx.Request().Query("orderBy")
	orderDir := ctx.Request().Query("orderDir")
	if orderBy != "" {
		return func(query orm.Query) orm.Query {
			return query.OrderBy(orderBy, orderDir)
		}
	}
	return func(query orm.Query) orm.Query {
		return query
	}
}

// ScpoeWith 关联查询
func (c *CrudService[Model]) ScpoeWith(ctx http.Context) func(orm.Query) orm.Query {
	withs := ctx.Request().Query("with")
	if withs != "" {
		return func(query orm.Query) orm.Query {
			for _, v := range strings.Split(withs, ",") {
				query = query.With(v)
			}
			return query
		}
	}
	return func(query orm.Query) orm.Query {
		return query
	}
}

// GetDetail 详情查询
func (c *CrudService[Model]) GetDetail(ctx http.Context, id any, scopes ...func(orm.Query) orm.Query) (item any, err error) {
	var data Model
	scopes = append(scopes, c.ScpoeOrderBy(ctx))
	query := c.Query(ctx).
		Scopes(scopes...).
		Where(consts.ID, id)
	err = query.FindOrFail(&data)
	return &data, err
}

func (c *CrudService[Model]) GetDetailScopes(ctx http.Context) []func(orm.Query) orm.Query {
	return []func(orm.Query) orm.Query{}
}

func (c *CrudService[Model]) Create(ctx http.Context, item any, scopes ...func(orm.Query) orm.Query) (err error) {
	tx, err := c.Query(ctx).BeginTransaction()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	query := c.Query(ctx).Scopes(scopes...)
	err = query.Create(item)
	if err != nil {
		return err
	}
	withs := ctx.Request().Query("with")
	if withs != "" {
		withList := strings.Split(withs, ",")
		for _, v := range withList {
			facades.Orm().WithContext(ctx).Query().Model(&item).Association(v).
				Append(reflect.ValueOf(item))
		}
	}
	return err
}
func (c *CrudService[Model]) GetCreateScopes(ctx http.Context) []func(orm.Query) orm.Query {
	return []func(orm.Query) orm.Query{}
}

func (c *CrudService[Model]) GetCreateData(ctx http.Context) any {
	var item Model
	ctx.Request().Bind(&item)
	return item
}

func (c *CrudService[Model]) Update(ctx http.Context, id any, item any, scopes ...func(orm.Query) orm.Query) (rowsAffected int64, err error) {

	tx, err := c.Query(ctx).BeginTransaction()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

	}()

	facades.Log().Info(ctx, fmt.Sprintf("Update %v with %v", id, item))
	query := c.Query(ctx).Scopes(scopes...)
	result, err := query.Where(consts.ID, id).Update(&item)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected, err
}

func (c *CrudService[Model]) GetUpdateScopes(ctx http.Context) []func(orm.Query) orm.Query {
	return []func(orm.Query) orm.Query{}
}

func (c *CrudService[Model]) Delete(ctx http.Context, id any, scopes ...func(orm.Query) orm.Query) (rowsAffected int64, err error) {
	query := c.Query(ctx).Scopes(scopes...)
	result, err := query.Where(consts.ID, id).Delete()
	return result.RowsAffected, err
}

func (c *CrudService[Model]) GetDeleteScopes(ctx http.Context) []func(orm.Query) orm.Query {
	return []func(orm.Query) orm.Query{}
}
