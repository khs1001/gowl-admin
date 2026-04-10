package core

import (
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type IndexReq struct {
	Page     int    `json:"page" `
	PerPage  int    `json:"perPage" `
	OrderBy  string `json:"orderBy" `
	OrderDir string `json:"orderDir" `
}

type IndexRes struct {
	Total int64 `json:"total" `
	Items any   `json:"items" `
}

type CrudController struct {
	*BaseController
	Model any
}

func (c *CrudController) Query(ctx http.Context) orm.Query {

	return facades.Orm().WithContext(ctx).Query().Model(c.Model)
}

// List 通用列表查询，支持分页和排序
// ctx: HTTP上下文
// scopes: 可选的查询作用域
// 返回：数据列表、总条数、错误信息
func (c *CrudController) List(ctx http.Context, scopes ...func(orm.Query) orm.Query) (items any, total int64, err error) {

	var req IndexReq
	items = NewModelSlice(c.Model)
	query := c.Query(ctx)
	err = gconv.Struct(ctx.Request().Queries(), &req)
	if err != nil {
		return nil, 0, err
	}
	if req.OrderBy != "" {
		query = query.OrderBy(req.OrderBy, req.OrderDir)
	}
	query = query.Scopes(scopes...)
	if req.Page > 0 {
		err = query.Paginate(req.Page, req.PerPage, items, &total)
	} else {
		err = query.Get(items)
	}
	if err != nil {
		return nil, 0, err
	}
	return items, total, err
}

func (c *CrudController) Index(ctx http.Context) http.Response {
	items, total, err := c.List(ctx)
	if err != nil {
		return Error(ctx, err)
	}
	return Success(ctx, &IndexRes{
		Total: total,
		Items: items,
	})
}

func (c *CrudController) Store(ctx http.Context) http.Response {
	item := ctx.Request().All()
	err := c.Query(ctx).Create(item)
	if err != nil {
		return Error(ctx, err)
	}
	return Success(ctx, item)
}

func (c *CrudController) Show(ctx http.Context) http.Response {
	var item any = NewModel(c.Model)
	id := ctx.Request().RouteInt64("id")
	err := c.Query(ctx).Where("id", id).FindOrFail(&item)
	if err != nil {
		return Error(ctx, err)
	}
	return Success(ctx, item)
}

func (c *CrudController) Update(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt64("id")
	result, err := c.Query(ctx).Where("id", id).Update(ctx.Request().All())
	if err != nil {
		return Error(ctx, err)
	}
	return Success(ctx, result)
}

func (c *CrudController) Destroy(ctx http.Context) http.Response {
	ids := ctx.Request().Route("id")
	result, err := c.Query(ctx).Where("id", strings.Split(ids, ",")).Delete()
	if err != nil {
		return Error(ctx, err)
	}
	return Success(ctx, result)
}
