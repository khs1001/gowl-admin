package core

import (
	"strings"

	"github.com/goravel/framework/contracts/http"
)

type ListParams struct {
	IsTreeList    bool
	ParentIdField string
	IdField       string
	ChildrenField string
	RootParentId  interface{}
}

type CrudController[T any] struct {
	*BaseController
	ListParams ListParams
	Service    ICrudService
}

func NewCrudController[Model any]() *CrudController[Model] {
	return &CrudController[Model]{
		Service: NewCrudService[Model](),
	}
}

func (c *CrudController[T]) Index(ctx http.Context) http.Response {
	var items any
	var total int64
	var err error
	scopes := c.Service.GetListScopes(ctx)
	if c.ListParams.IsTreeList {
		var data any
		data, err = c.Service.GetList(ctx, scopes...)
		items = ListToTree(data,
			c.ListParams.ParentIdField,
			c.ListParams.IdField,
			c.ListParams.ChildrenField,
			c.ListParams.RootParentId)
	} else {
		page := ctx.Request().InputInt("page", 0)
		perPage := ctx.Request().InputInt("perPage", 0)
		items, total, err = c.Service.GetPageList(ctx, page, perPage, scopes...)
	}
	if err != nil {
		return c.Error(ctx, err)
	}
	return c.Success(ctx, http.Json{
		"items": items,
		"total": total,
	})
}
func (c *CrudController[T]) Show(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")
	scopes := c.Service.GetDetailScopes(ctx)
	item, err := c.Service.GetDetail(ctx, id, scopes...)
	if err != nil {
		return c.Error(ctx, err)
	}
	return c.Success(ctx, item)
}
func (c *CrudController[T]) Store(ctx http.Context) http.Response {
	item := c.Service.GetCreateData(ctx)
	scopes := c.Service.GetCreateScopes(ctx)
	err := c.Service.Create(ctx, item, scopes...)
	if err != nil {
		return c.Error(ctx, err)
	}
	return c.Success(ctx, http.Json{})
}

func (c *CrudController[T]) Update(ctx http.Context) http.Response {
	var item T
	ctx.Request().Bind(&item)
	id := ctx.Request().RouteInt("id")
	scopes := c.Service.GetUpdateScopes(ctx)
	rowsAffected, err := c.Service.Update(ctx, id, item, scopes...)
	if err != nil {
		return c.Error(ctx, err)
	}
	return c.Success(ctx, http.Json{
		"rows_affected": rowsAffected,
	})
}

func (c *CrudController[T]) Destroy(ctx http.Context) http.Response {
	ids := strings.Split(ctx.Request().Route("id"), ",")
	scopes := c.Service.GetDeleteScopes(ctx)
	rowsAffected, err := c.Service.Delete(ctx, ids, scopes...)
	if err != nil {
		return c.Error(ctx, err)
	}
	return c.Success(ctx, http.Json{
		"rows_affected": rowsAffected,
	})
}
