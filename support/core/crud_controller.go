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

	items, total, err = c.Service.GetList(ctx)
	if c.ListParams.IsTreeList {
		items = ListToTree(items,
			c.ListParams.ParentIdField,
			c.ListParams.IdField,
			c.ListParams.ChildrenField,
			c.ListParams.RootParentId)
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
	item, err := c.Service.GetDetail(ctx, id)
	if err != nil {
		return c.Error(ctx, err)
	}
	return c.Success(ctx, item)
}

func (c *CrudController[T]) Store(ctx http.Context) http.Response {
	var item T
	err := ctx.Request().Bind(&item)
	if err != nil {
		return c.Error(ctx, err)
	}
	err = c.Service.Create(ctx, &item)
	if err != nil {
		return c.Error(ctx, err)
	}
	return c.Success(ctx, &item)
}

func (c *CrudController[T]) Update(ctx http.Context) http.Response {
	var item T
	err := ctx.Request().Bind(&item)
	if err != nil {
		return c.Error(ctx, err)
	}
	id := ctx.Request().RouteInt("id")
	rowsAffected, err := c.Service.Update(ctx, id, &item)
	if err != nil {
		return c.Error(ctx, err)
	}
	return c.Success(ctx, http.Json{
		"rows_affected": rowsAffected,
	})
}

func (c *CrudController[T]) Destroy(ctx http.Context) http.Response {
	var err error
	ids := strings.Split(ctx.Request().Route("id"), ",")
	rowsAffected, err := c.Service.Delete(ctx, ids)
	if err != nil {
		return c.Error(ctx, err)
	}
	return c.Success(ctx, http.Json{
		"rows_affected": rowsAffected,
	})
}
