package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/khs1001/gowl-admin/consts"
	"github.com/khs1001/gowl-admin/http/responses"
	"github.com/khs1001/gowl-admin/models"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	s := &AuthService{}
	return s
}

func (s *AuthService) Login(ctx http.Context, username, password string) (token string, err error) {

	var user models.AdminUser
	err = facades.Orm().WithContext(ctx).Query().Where(consts.Username, username).FirstOrFail(&user)
	if err != nil || !facades.Hash().Check(password, user.Password) {
		return "", errors.New(consts.ErrorLoginFailed)
	}
	if user.Enabled != consts.Enabled {
		return "", errors.New(consts.ErrorUserDisabled)
	}
	// 设置token
	guard := facades.Config().GetString(consts.AdminAuthGuard)
	return facades.Auth(ctx).Guard(guard).LoginUsingID(user.ID)
}

func (s *AuthService) BuildRoutes(ctx context.Context, items []*models.AdminMenu, parentID uint) []*responses.AdminRoute {
	var result []*responses.AdminRoute
	var idStr string

	var _component string
	for _, v := range items {
		if uint(v.ParentId) == parentID {
			idStr = fmt.Sprintf("[%d]", v.ID)
			pageSign := ""
			switch v.UrlType {
			case 3:
				_component = "iframe"
			case 4:
				_component = "amis"
				pageSign = gstr.Split(v.Url, "?")[0] + "/index"
			default:
				_component = v.Component
			}
			if _component == "" {
				_component = "amis"
				pageSign = gstr.Split(v.Url, "?")[0] + "/index"
			}
			item := &responses.AdminRoute{
				// 修改 Name 为 Path 的 MD5 值
				Name:      idStr,
				Path:      v.Url,
				Component: _component,
				IsHome:    int(v.IsHome),
				IframeUrl: v.IframeUrl,
				UrlType:   int(v.UrlType),
				KeepAlive: int(v.KeepAlive),
				IsFull:    int(v.IsFull),
				IsLink:    v.UrlType == 2,
				PageSign:  pageSign,
				Meta: &responses.AdminRouteMeta{
					Title:       v.Title,
					Icon:        v.Icon,
					CustomOrder: v.CustomOrder,
					Hide:        v.Visible == 0,
				},
			}
			children := s.BuildRoutes(ctx, items, v.ID)
			if len(children) > 0 {
				item.Children = children
				item.Component = _component
			}
			result = append(result, item)
		}
	}
	return result
}

func (s *AuthService) List2Tree(items []*models.AdminMenu, parentID uint) (result []*models.AdminMenu) {
	if len(items) == 0 {
		return nil
	}
	result = make([]*models.AdminMenu, 0)
	for _, item := range items {
		if item.ParentId == int(parentID) {
			children := s.List2Tree(items, item.ID)
			if children != nil {
				item.Children = children
			}
			result = append(result, item)
		}
	}
	return result
}
