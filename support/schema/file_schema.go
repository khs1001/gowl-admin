package schema

import (
	"bytes"
	"context"
	"encoding/json"
	"html/template"
	"path/filepath"
	"regexp"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/file"
)

const (
	NodeExp     = "NodeExp"
	TemplateExp = "TemplateExp"
)

// FileSchema 定义文件模式的结构体
type FileSchema struct {
	Data map[string]any // 数据
	Sign string         // 标识
	Path string         // 路径
}

// SchemaRoot 定义模式文件的根目录
var SchemaRoot = "resources/schemas"

// GetSchema 获取指定标识和路径的模式
// ctx: 上下文
// 返回: 模式的 JSON 对象
func (s *FileSchema) GetSchema(ctx context.Context) *gjson.Json {
	if s.Data == nil {
		s.Data = make(map[string]any)
	}
	s.Data["Path"] = s.Path
	facades.Log().Info("Data ", s.Data)
	schema, err := s.GetfileSchema(ctx, s.Path, s.Sign)
	if err != nil {
		return nil
	}
	return schema
}

func (s *FileSchema) GetfileSchema(ctx context.Context, path, sign string) (schema *gjson.Json, err error) {
	targetFile := sign + ".json"
	currentDir := path
	for {
		// 构建完整文件路径
		schemaFile := filepath.Join(SchemaRoot, currentDir, targetFile)
		// 检查文件是否存在
		if file.Exists(schemaFile) {
			content, err := file.GetContent(schemaFile)
			if err != nil {
				continue
			}
			schema = gjson.New(ReplaceUserSettingTag(content))
			if !schema.IsNil() {
				s.ScanNode(ctx, schema, currentDir, sign)
				break
			}
		}
		if currentDir == filepath.Dir(currentDir) {
			break
		}
		currentDir = filepath.Dir(currentDir)
	}
	return schema, nil
}

func (s *FileSchema) ScanNode(ctx context.Context, parent *gjson.Json, path, sign string) *gjson.Json {
	switch nv := parent.Var().Interface().(type) {
	case map[string]interface{}:
		for key, value := range parent.GetJsonMap(".") {
			newJson := s.ScanNode(ctx, value, path, sign)
			parent.Set(key, newJson.Var().Interface())
		}
		//继承的节点
		s.ExtendNode(ctx, parent, path, sign)
	case []interface{}:
		for key, value := range parent.GetJsons(".") {
			newJson := s.ScanNode(ctx, value, path, sign)
			parent.Set(gconv.String(key), newJson.Var().Interface())
		}
	case string:
		//是否包括语法
		if gstr.StrTill(gstr.StrEx(nv, "{{"), "}}") != "" {
			newJson := s.ParseTemplate(ctx, nv, path, sign)
			parent = s.ScanNode(ctx, newJson, path, sign)
		}
	}
	return parent
}

func (s *FileSchema) ExtendNode(ctx context.Context, parent *gjson.Json, path, sign string) {
	if !parent.Get("type").IsMap() {
		return
	}
	typeSchema := parent.GetJson("type")
	//移除原type
	parent.Remove("type")
	for key, value := range typeSchema.Map() {
		if parent.Contains(key) {
			continue
		}
		parent.Set(key, value)
	}
}

func ReplaceUserSettingTag(input string) string {
	// 定义正则表达式匹配模式
	pattern := `"@([a-zA-Z0-9_#.]+)"`
	// 编译正则表达式
	re := regexp.MustCompile(pattern)
	// 替换匹配的字符串
	return re.ReplaceAllString(input, `"{{node \"$1\"}}"`)
}

func (s *FileSchema) getNodeSchema(ctx context.Context, path string, currSign string) template.HTML {
	node, err := s.GetfileSchema(ctx, path, currSign)
	if err != nil {
		return template.HTML("null")
	}
	nodeStr, err := json.Marshal(node)
	if err != nil {
		return template.HTML("null")
	}
	if nodeStr == nil {
		return template.HTML("null")
	}
	return template.HTML(string(nodeStr))
}

func (s *FileSchema) ParseTemplate(ctx context.Context, content string, path, sign string) (schema *gjson.Json) {
	content = ReplaceUserSettingTag(content)
	customFuncMap := funcMap
	customFuncMap["node"] = func(currSign string) template.HTML {
		currPath := s.Path
		if currSign == sign {
			currPath = filepath.Dir(currPath)
		}
		return s.getNodeSchema(ctx, currPath, currSign)
	}
	templateObject := template.New("amis")
	templateObject.Funcs(customFuncMap)
	tmpl, err := templateObject.Parse(content)
	if err != nil {
		return nil
	}
	var result bytes.Buffer
	err = tmpl.Execute(&result, s.Data)
	if err != nil {
		return nil
	}
	return gjson.New(result.String())
}
