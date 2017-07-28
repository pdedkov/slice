package slice_test

import (
	"fmt"
	"github.com/pdedkov/slice"
)

func ExampleFilter() {
	v := []string{"", "1", "2"}

	ret, _ := slice.Filter(v, func(val interface{}) bool {
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
	})
	fmt.Printf("slice len is %d", len(ret))

	// Output:
	// slice len is 2
}
