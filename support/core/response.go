package core

import (
	"github.com/goravel/framework/contracts/http"
)

type ResponseData struct {
	Status            int         `json:"status"`
	Msg               string      `json:"msg"`
	DoNotDisplayToast int         `json:"doNotDisplayToast"`
	Data              interface{} `json:"data"`
}

// Success 响应成功
func Success(ctx http.Context, data any) http.Response {
	return ctx.Response().Success().Json(
		&ResponseData{
			Data:              data,
			DoNotDisplayToast: 1,
		})
}

// Ok 响应成功消息
func Ok(ctx http.Context, msg string) http.Response {
	return ctx.Response().Success().Json(
		&ResponseData{
			Msg:               msg,
			DoNotDisplayToast: 0,
		})
}

// Fail 响应失败消息，弹出提示
func Fail(ctx http.Context, data any) http.Response {
	return ctx.Response().Success().Json(
		&ResponseData{
			Status:            500,
			Data:              data,
			DoNotDisplayToast: 0,
		})
}

// Error 响应失败消息
func Error(ctx http.Context, err error) http.Response {
	return ctx.Response().Success().Json(
		&ResponseData{
			Status:            500,
			Msg:               err.Error(),
			DoNotDisplayToast: 1,
		})
}

// UnAuthorized 未登录
func UnAuthorized(ctx http.Context) {
	ctx.Request().AbortWithStatusJson(http.StatusOK,
		http.Json{
			"code":              401,
			"msg":               "",
			"doNotDisplayToast": 1,
		})
}

func Forbidden(ctx http.Context) {
	ctx.Request().AbortWithStatusJson(http.StatusOK,
		&ResponseData{
			Status:            http.StatusForbidden,
			Msg:               "没有权限",
			DoNotDisplayToast: 0,
		})
}

func AbortOnError(ctx http.Context, err error) {
	if err != nil {
		ctx.Request().AbortWithStatusJson(http.StatusOK,
			&ResponseData{
				Status:            http.StatusInternalServerError,
				Msg:               err.Error(),
				DoNotDisplayToast: 0,
			})
	}

}
