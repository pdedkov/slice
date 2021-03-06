package slice_test

import (
	"github.com/pdedkov/slice"
	"testing"
)

func TestFilter(t *testing.T) {
	val := []string{"", "1", "2"}
	fn := func(val interface{}) bool {
		switch val.(type) {
		case string:
			if len(val.(string)) > 0 {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	}

	arr, _ := slice.Filter(val, fn)
	if len(arr) != 2 {
		t.Errorf("fail %+v\n", arr)
	}
}
