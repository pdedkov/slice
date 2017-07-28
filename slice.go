package slice

import (
	"errors"
	"reflect"
)

// Filter filter slice with filter function
func Filter(s interface{}, fn func(interface{}) bool) ([]interface{}, error) {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		return nil, errors.New("Invalid argument. Given non-slice type argument")
	}
	var r []interface{}
	for i := 0; i < val.Len(); i++ {
		if fn(val.Index(i).Interface()) {
			r = append(r, val.Index(i).Interface())
		}
	}
	return r, nil
}
