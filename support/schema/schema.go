package schema

import (
	"context"
	"html/template"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/text/gstr"
)

const (
	QueryPath = "QueryPath"
)

var funcMap = make(template.FuncMap)

// 注册模板函数
func RegisterFunc(name string, f interface{}) {
	funcMap[name] = f
}

func GetPage(ctx context.Context, sign string, path string, data map[string]any) *gjson.Json {
	path = gstr.Split(path, "?")[0]
	s := &FileSchema{
		Sign: sign,
		Path: path,
		Data: data,
	}
	return gjson.New(s.GetSchema(ctx))
}
