package core

import (
	"reflect"
)

func NewModel(tableNameOrStruct any) any {
	// 获取反射类型
	t := reflect.TypeOf(tableNameOrStruct)

	// 检查输入是否为指针类型
	if t.Kind() != reflect.Ptr {
		panic("输入必须是指针类型")
	}

	// 获取指针指向的元素类型（结构体类型）
	structType := t.Elem()
	// 创建一个新的实例
	newInstance := reflect.New(structType).Interface()
	return newInstance
}

func NewModelSlice(tableNameOrStruct any) any {
	// 获取反射类型
	t := reflect.TypeOf(tableNameOrStruct)
	// 获取指针指向的元素类型（结构体类型）
	structType := t.Elem()
	// 创建结构体指针类型的切片（指针数组）
	// 先获取结构体指针类型，再创建该类型的切片
	ptrType := reflect.PointerTo(structType)
	sliceType := reflect.SliceOf(ptrType)
	slice := reflect.MakeSlice(sliceType, 0, 0)

	return slice.Interface()
}
