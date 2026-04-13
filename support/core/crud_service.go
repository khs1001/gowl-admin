package core

import (
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/khs1001/gowl-admin/consts"
)

type ICrudService interface {
	GetList(ctx http.Context, scopes ...func(orm.Query) orm.Query) (items any, total int64, err error)
	GetDetail(ctx http.Context, id any, scopes ...func(orm.Query) orm.Query) (item any, err error)
	Create(ctx http.Context, item any, scopes ...func(orm.Query) orm.Query) (err error)
	Update(ctx http.Context, id any, item any, scopes ...func(orm.Query) orm.Query) (rowsAffected int64, err error)
	Delete(ctx http.Context, id any, scopes ...func(orm.Query) orm.Query) (rowsAffected int64, err error)
	SetAssociations(associations []string)
	SetWiths(withs []string)
}

// CrudService 通用增删改查服务
type CrudService[Model any] struct {
	Withs        []string
	Associations []string
}

func NewCrudService[Model any]() *CrudService[Model] {
	return &CrudService[Model]{}
}

func (c *CrudService[Model]) SetWiths(withs []string) {
	c.Withs = withs
}

func (c *CrudService[Model]) SetAssociations(associations []string) {
	c.Associations = associations
}

func (s *CrudService[Model]) GetList(ctx http.Context, scopes ...func(orm.Query) orm.Query) (items any, total int64, err error) {
	var model Model
	items = make([]*Model, 0)
	//增加排序
	scopes = append(scopes, s.ScpoeWith(ctx), s.ScpoeOrderBy(ctx))
	page := ctx.Request().InputInt("page", 0)
	perPage := ctx.Request().InputInt("perPage", 0)
	query := facades.Orm().WithContext(ctx).Query().Model(&model)
	if page > 0 && perPage > 0 {
		err = query.Scopes(scopes...).Paginate(page, perPage, &items, &total)
	} else {
		err = query.Scopes(scopes...).Get(&items)
	}
	return items, total, err
}

// ScpoeOrderBy 排序
func (s *CrudService[Model]) ScpoeOrderBy(ctx http.Context) func(orm.Query) orm.Query {
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
func (s *CrudService[Model]) ScpoeWith(ctx http.Context) func(orm.Query) orm.Query {
	return func(query orm.Query) orm.Query {
		for _, v := range s.Withs {
			query = query.With(v)
		}
		return query
	}
}

// GetDetail 详情查询
func (s *CrudService[Model]) GetDetail(ctx http.Context, id any, scopes ...func(orm.Query) orm.Query) (item any, err error) {
	var data Model

	query := facades.Orm().WithContext(ctx).Query().Model(&data)
	scopes = append(scopes, s.ScpoeOrderBy(ctx))
	err = query.
		Scopes(scopes...).
		Where(consts.ID, id).
		FindOrFail(&data)
	return &data, err
}

func (s *CrudService[Model]) Create(ctx http.Context, item any, scopes ...func(orm.Query) orm.Query) (err error) {
	tx, err := facades.Orm().WithContext(ctx).Query().BeginTransaction()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Create(item)
	if err != nil {
		return err
	}
	for _, v := range s.Associations {
		value, has := GetStructField(item, v)
		if !has {
			continue
		}
		err = tx.Model(item).Association(v).Append(value)
		if err != nil {
			break
		}
	}
	return err
}

func (s *CrudService[Model]) Update(ctx http.Context, id any, item any, scopes ...func(orm.Query) orm.Query) (rowsAffected int64, err error) {
	tx, err := facades.Orm().WithContext(ctx).Query().BeginTransaction()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	for _, v := range s.Associations {
		value, has := GetStructField(item, v)
		if !has {
			continue
		}
		err = tx.Model(item).Association(v).Replace(value)
		if err != nil {
			break
		}
	}
	result, err := tx.Model(item).Where(consts.ID, id).Update(item)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected, err
}

func (s *CrudService[Model]) Delete(ctx http.Context, id any, scopes ...func(orm.Query) orm.Query) (rowsAffected int64, err error) {
	tx, err := facades.Orm().WithContext(ctx).Query().BeginTransaction()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	items := make([]*Model, 0)
	err = tx.Model(items).Where(consts.ID, id).Get(&items)
	for _, item := range items {
		for _, v := range s.Associations {
			err = tx.Model(item).Association(v).Clear()
			if err != nil {
				return rowsAffected, err
			}
		}
		result, err := tx.Delete(item)
		if err != nil {
			return rowsAffected, err
		}
		rowsAffected += result.RowsAffected
	}
	return rowsAffected, err
}
