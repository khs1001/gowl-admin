package core

import "github.com/goravel/framework/contracts/http"

type BaseController struct {
}

// Success 响应成功
func (c *BaseController) Success(ctx http.Context, data any) http.Response {
	return ctx.Response().Success().Json(
		&ResponseData{
			Data:              data,
			DoNotDisplayToast: 1,
		})
}

// Error 响应失败消息
func (c *BaseController) Error(ctx http.Context, err error) http.Response {
	return ctx.Response().Success().Json(
		&ResponseData{
			Status:            500,
			Msg:               err.Error(),
			DoNotDisplayToast: 1,
		})
}
