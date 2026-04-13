package core

import "reflect"

// GetStructField 获取 any 类型变量指向的结构体的指定字段值
// 参数：
//   obj: any 类型（结构体/结构体指针）
//   fieldName: 字段名（必须首字母大写，可导出）
// 返回：字段值、是否存在
func GetStructField(obj any, fieldName string) (any, bool) {
	// 1. 获取反射值，自动解指针（支持结构体指针、结构体本身）
	val := reflect.ValueOf(obj)
	for val.Kind() == reflect.Ptr {
		val = val.Elem() // 解指针
	}

	// 2. 校验是否是结构体
	if val.Kind() != reflect.Struct {
		return nil, false
	}

	// 3. 获取字段
	field := val.FieldByName(fieldName)
	// 4. 校验字段是否存在且可导出
	if !field.IsValid() || !field.CanInterface() {
		return nil, false
	}

	// 5. 返回字段值
	return field.Interface(), true
}
