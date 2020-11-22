package tools

import (
	"reflect"
)

// IsContains 查找值val是否在数组array中存在
func IsContains(val interface{}, array interface{}) bool {
	if array == nil {
		return false
	}
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				return true
			}
		}
	}
	return false
}
