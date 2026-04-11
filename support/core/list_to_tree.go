package core

import (
	"reflect"

	"github.com/gogf/gf/v2/util/gconv"
)

// ListToTree 利用反射将列表转换为树形结构
// nodes: 输入的节点列表（必须是切片类型）
// parentIdField: 父ID字段名
// idField: 当前节点ID字段名
// childrenField: 子节点字段名
// rootParentId: 根节点的父ID值
func ListToTree(nodes interface{}, parentIdField, idField, childrenField string, rootParentId interface{}) interface{} {
	// 检查输入是否为切片
	val := reflect.ValueOf(nodes)
	if val.Kind() != reflect.Slice {
		return nil
	}
	// 创建结果切片（与输入切片相同类型）
	result := reflect.MakeSlice(val.Type(), 0, 0)
	// 遍历所有节点
	for i := 0; i < val.Len(); i++ {
		node := val.Index(i)
		// 获取节点的父ID值
		parentIdVal := GetFieldValue(node, parentIdField)
		if gconv.String(parentIdVal) == gconv.String(rootParentId) {
			// 递归查找子节点
			children := ListToTree(nodes, parentIdField, idField, childrenField, GetFieldValue(node, idField))
			// 设置子节点
			if children != nil {
				SetFieldValue(node, childrenField, children)
			}

			// 添加到结果集
			result = reflect.Append(result, node)
		}
	}

	return result.Interface()
}

// 获取字段值
func GetFieldValue(node reflect.Value, fieldName string) any {
	// 如果是指针，获取其指向的值
	if node.Kind() == reflect.Ptr {
		node = node.Elem()
	}

	// 检查是否为结构体
	if node.Kind() != reflect.Struct {
		return reflect.Value{}.Interface()
	}
	field := node.FieldByName(fieldName)
	if !field.IsValid() || !field.CanSet() {
		return reflect.Value{}.Interface()
	}
	// 获取字段值
	return field.Interface()
}

// 设置字段值
func SetFieldValue(node reflect.Value, fieldName string, value interface{}) {
	// 如果是指针，获取其指向的值
	if node.Kind() == reflect.Ptr {
		node = node.Elem()
	}

	// 检查是否为结构体
	if node.Kind() != reflect.Struct {
		return
	}

	// 获取字段
	field := node.FieldByName(fieldName)
	if !field.IsValid() || !field.CanSet() {
		return
	}

	// 设置字段值
	field.Set(reflect.ValueOf(value))
}
