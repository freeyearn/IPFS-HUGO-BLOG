package structs

import (
	"reflect"
)

func StructAssign(binding interface{}, value interface{}, tagName string) {
	bVal := reflect.ValueOf(binding)
	if bVal.Kind() == reflect.Ptr {
		bVal = bVal.Elem()
	} //获取reflect.Type类型
	vVal := reflect.ValueOf(value)
	if vVal.Kind() == reflect.Ptr {
		vVal = vVal.Elem()
	} //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	bTypeOfT := bVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		if !vVal.Field(i).IsZero() {
			vName := vTypeOfT.Field(i).Tag.Get(tagName)
			for j := 0; j < bVal.NumField(); j++ {
				bName := bTypeOfT.Field(j).Tag.Get(tagName)
				if vName == bName {
					bVal.Field(j).Set(reflect.ValueOf(vVal.Field(i).Interface()))
					break
				}
			}
		}
	}
}
